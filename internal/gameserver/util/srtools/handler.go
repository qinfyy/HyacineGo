package srtools

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Dir        string
	CorsOrigin string
}

type Handler struct {
	opts Options
}

func New(opts Options) *Handler {
	if opts.CorsOrigin == "" {
		opts.CorsOrigin = "*"
	}
	return &Handler{opts: opts}
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	_ = r
	h.addCorsHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	h.addCorsHeaders(w)

	status := http.StatusOK
	message := "OK"

	body, _ := io.ReadAll(io.LimitReader(r.Body, 10<<20))
	if len(body) != 0 {
		var parsed any
		if err := json.Unmarshal(body, &parsed); err != nil {
			status = http.StatusBadRequest
			message = "invalid JSON payload"
		} else {
			obj, ok := parsed.(map[string]any)
			if ok {
				name := "freesr-data.json"
				if v, ok := obj["name"]; ok {
					if s, ok := v.(string); ok && s != "" {
						name = s
					}
				}
				if v, ok := obj["data"]; ok {
					if _, ok := v.(map[string]any); ok {
						if err := h.saveFreesrData(name, v); err != nil {
							status = http.StatusInternalServerError
							message = err.Error()
						}
					} else {
						status = http.StatusBadRequest
						message = "srtools data must be an object"
					}
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status":  status,
		"message": message,
	})
}

func (h *Handler) saveFreesrData(name string, v any) error {
	dir := h.opts.Dir
	if dir == "" {
		return nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	name = filepath.Base(name)
	if name == "." || name == string(filepath.Separator) || name == "" {
		name = "freesr-data.json"
	}
	if !strings.HasSuffix(strings.ToLower(name), ".json") {
		name += ".json"
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, name)
	return os.WriteFile(path, raw, 0o644)
}

func (h *Handler) addCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", h.opts.CorsOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Max-Age", "86400")
}
