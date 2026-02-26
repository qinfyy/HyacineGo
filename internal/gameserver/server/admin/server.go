package admin

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"time"

	observability "hyacine-server/internal/common"
)

type Options struct {
	Addr       string
	ConfigPath string
	Reload     func() error
	Metrics    *observability.Metrics
}

type Server struct {
	opts Options
	srv  *http.Server
	ln   net.Listener
}

func New(opts Options) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if opts.Reload == nil {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		if err := opts.Reload(); err != nil {
			slog.Error("reload failed", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("reloaded\n"))
	})
	if opts.Metrics != nil {
		mux.Handle("/metrics", opts.Metrics.Handler())
	}

	return &Server{
		opts: opts,
		srv: &http.Server{
			Addr:              opts.Addr,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
		},
	}
}

func (s *Server) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	go func() {
		<-ctx.Done()
		_ = s.Shutdown(context.Background())
	}()

	err = s.srv.Serve(ln)
	if err == nil || errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}
	return s.srv.Shutdown(ctx)
}
