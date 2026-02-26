package observability

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"strings"
	"time"
)

// prettyHandler is a small colored console handler similar to SR-CasPS-main's Zig logger output.
// It is intentionally simple and optimized for developer readability.
type prettyHandler struct {
	out   io.Writer
	level slog.Leveler
}

func newPrettyHandler(out io.Writer, level slog.Leveler) slog.Handler {
	return &prettyHandler{out: out, level: level}
}

func (h *prettyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

func (h *prettyHandler) Handle(_ context.Context, r slog.Record) error {
	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}

	level := r.Level.String()
	levelColored := colorLevel(r.Level, level)

	attrs := make([]slog.Attr, 0, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})
	sort.Slice(attrs, func(i, j int) bool { return attrs[i].Key < attrs[j].Key })

	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "%s %s %s", ts.Format("15:04:05.000"), levelColored, r.Message)
	for _, a := range attrs {
		_, _ = fmt.Fprintf(&b, " %s=%v", a.Key, a.Value.Any())
	}
	b.WriteByte('\n')
	_, err := io.WriteString(h.out, b.String())
	return err
}

func (h *prettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Keep it simple: materialize attrs as a prefix by wrapping WithGroup.
	// For our usage we rarely rely on WithAttrs.
	return &prettyHandlerWithAttrs{base: h, attrs: append([]slog.Attr(nil), attrs...)}
}

func (h *prettyHandler) WithGroup(name string) slog.Handler {
	if strings.TrimSpace(name) == "" {
		return h
	}
	return &prettyHandlerWithGroup{base: h, group: name}
}

type prettyHandlerWithAttrs struct {
	base  *prettyHandler
	attrs []slog.Attr
}

func (h *prettyHandlerWithAttrs) Enabled(ctx context.Context, level slog.Level) bool {
	return h.base.Enabled(ctx, level)
}
func (h *prettyHandlerWithAttrs) WithAttrs(attrs []slog.Attr) slog.Handler {
	merged := append(append([]slog.Attr(nil), h.attrs...), attrs...)
	return &prettyHandlerWithAttrs{base: h.base, attrs: merged}
}
func (h *prettyHandlerWithAttrs) WithGroup(name string) slog.Handler { return h.base.WithGroup(name) }
func (h *prettyHandlerWithAttrs) Handle(ctx context.Context, r slog.Record) error {
	r2 := r.Clone()
	for _, a := range h.attrs {
		r2.AddAttrs(a)
	}
	return h.base.Handle(ctx, r2)
}

type prettyHandlerWithGroup struct {
	base  *prettyHandler
	group string
}

func (h *prettyHandlerWithGroup) Enabled(ctx context.Context, level slog.Level) bool {
	return h.base.Enabled(ctx, level)
}
func (h *prettyHandlerWithGroup) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.base.WithAttrs(attrs)
}
func (h *prettyHandlerWithGroup) WithGroup(name string) slog.Handler {
	if strings.TrimSpace(name) == "" {
		return h
	}
	return &prettyHandlerWithGroup{base: h.base, group: h.group + "." + name}
}
func (h *prettyHandlerWithGroup) Handle(ctx context.Context, r slog.Record) error {
	r2 := r.Clone()
	// Prefix keys to mimic grouped fields.
	r.Attrs(func(a slog.Attr) bool {
		r2.AddAttrs(slog.Any(h.group+"."+a.Key, a.Value.Any()))
		return true
	})
	return h.base.Handle(ctx, r2)
}

func colorLevel(level slog.Level, s string) string {
	// Basic ANSI colors; safe to show even on Windows Terminal.
	switch {
	case level >= slog.LevelError:
		return "\x1b[31m" + s + "\x1b[0m" // red
	case level >= slog.LevelWarn:
		return "\x1b[33m" + s + "\x1b[0m" // yellow
	case level >= slog.LevelInfo:
		return "\x1b[36m" + s + "\x1b[0m" // cyan
	default:
		return "\x1b[90m" + s + "\x1b[0m" // gray
	}
}
