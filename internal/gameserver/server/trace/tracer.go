package trace

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Event struct {
	Time  string `json:"time"`
	Dir   string `json:"dir"` // RECV|SEND|UNHANDLED|AUTOREPLY
	CmdID uint16 `json:"cmd_id"`
	Cmd   string `json:"cmd"`
	Known bool   `json:"known"`
	Bytes int    `json:"bytes"`
	State string `json:"state,omitempty"`
	Note  string `json:"note,omitempty"`
}

type Tracer struct {
	mu   sync.Mutex
	path string
	f    *os.File
	enc  *json.Encoder
}

func New(path string) (*Tracer, error) {
	if path == "" {
		return &Tracer{}, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	// Single-file traces: always truncate on start to avoid keeping old traces.
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	return &Tracer{path: path, f: f, enc: json.NewEncoder(f)}, nil
}

func (t *Tracer) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.f == nil {
		return nil
	}
	err := t.f.Close()
	t.f = nil
	t.enc = nil
	return err
}

func (t *Tracer) Write(e Event) {
	if t == nil || t.enc == nil {
		return
	}
	e.Time = time.Now().Format(time.RFC3339Nano)
	t.mu.Lock()
	defer t.mu.Unlock()
	_ = t.enc.Encode(e)
}
