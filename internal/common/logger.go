package observability

import (
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"sync/atomic"
)

var currentLevel atomic.Value  // slog.Level
var currentLogger atomic.Value // *slog.Logger
var closeLogFile atomic.Value  // func()

func NewLogger(w io.Writer, level string) *slog.Logger {
	SetLevel(level)

	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: levelVar{},
	})
	l := slog.New(handler)
	currentLogger.Store(l)
	return l
}

type LoggerOptions struct {
	Level  string
	Format string // json|text|pretty
	File   string // optional
}

func ConfigureLogger(opts LoggerOptions) (*slog.Logger, func(), error) {
	SetLevel(opts.Level)

	var out io.Writer = os.Stdout
	var closer func()
	if strings.TrimSpace(opts.File) != "" {
		// Single-file logs: always truncate on start to avoid keeping old logs.
		f, err := os.OpenFile(opts.File, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return nil, nil, err
		}
		out = io.MultiWriter(os.Stdout, f)
		closer = func() { _ = f.Close() }
	}

	var handler slog.Handler
	switch strings.ToLower(strings.TrimSpace(opts.Format)) {
	case "pretty":
		handler = newPrettyHandler(out, levelVar{})
	case "text":
		handler = slog.NewTextHandler(out, &slog.HandlerOptions{Level: levelVar{}})
	default:
		handler = slog.NewJSONHandler(out, &slog.HandlerOptions{Level: levelVar{}})
	}

	l := slog.New(handler)
	currentLogger.Store(l)

	if c := closer; c != nil {
		closeLogFile.Store(c)
	}
	return l, closer, nil
}

func SetLevel(level string) {
	var lv slog.Level
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		lv = slog.LevelDebug
	case "warn", "warning":
		lv = slog.LevelWarn
	case "error":
		lv = slog.LevelError
	default:
		lv = slog.LevelInfo
	}
	currentLevel.Store(lv)
}

type levelVar struct{}

func (levelVar) Level() slog.Level {
	if v := currentLevel.Load(); v != nil {
		return v.(slog.Level)
	}
	return slog.LevelInfo
}

var once sync.Once

func Default() *slog.Logger {
	if v := currentLogger.Load(); v != nil {
		return v.(*slog.Logger)
	}
	once.Do(func() {
		l := NewLogger(os.Stdout, "info")
		slog.SetDefault(l)
	})
	if v := currentLogger.Load(); v != nil {
		return v.(*slog.Logger)
	}
	return slog.Default()
}
