package command

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"hyacine-server/internal/gameserver/database"
)

type Context struct {
	Out      io.Writer
	DataBase *database.DataBase
}

type Handler func(ctx Context, args []string) error

type Registry struct {
	handlers map[string]Handler
}

func NewRegistry() *Registry {
	return &Registry{handlers: map[string]Handler{}}
}

func (r *Registry) Register(name string, h Handler) {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" || h == nil {
		return
	}
	r.handlers[name] = h
}

func (r *Registry) ExecuteLine(ctx Context, line string) error {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	line = strings.TrimPrefix(line, "/")
	fields := splitFields(line)
	if len(fields) == 0 {
		return nil
	}
	cmd := strings.ToLower(fields[0])
	args := fields[1:]

	h, ok := r.handlers[cmd]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd)
	}
	return h(ctx, args)
}

func splitFields(s string) []string {
	// Allows quoted args: /account create "a@b.com"
	var out []string
	r := bufio.NewReader(strings.NewReader(s))
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			break
		}
		for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			ch, _, err = r.ReadRune()
			if err != nil {
				return out
			}
		}
		var token strings.Builder
		if ch == '"' {
			for {
				ch, _, err = r.ReadRune()
				if err != nil || ch == '"' {
					break
				}
				token.WriteRune(ch)
			}
		} else {
			token.WriteRune(ch)
			for {
				ch, _, err = r.ReadRune()
				if err != nil {
					break
				}
				if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
					break
				}
				token.WriteRune(ch)
			}
		}
		out = append(out, token.String())
		if err != nil {
			break
		}
	}
	return out
}
