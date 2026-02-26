package program

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"hyacine-server/internal/gameserver/command"
	"hyacine-server/internal/gameserver/database"
)

type Service interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Program struct {
	services []Service
	db       *database.DataBase
	commands *command.Registry
	out      io.Writer
}

type Options struct {
	Services []Service
	DataBase *database.DataBase
	Out      io.Writer
}

func New(opts Options) *Program {
	out := opts.Out
	if out == nil {
		out = os.Stdout
	}
	p := &Program{
		services: opts.Services,
		db:       opts.DataBase,
		commands: command.NewRegistry(),
		out:      out,
	}
	command.RegisterAccount(p.commands)
	return p
}

func (p *Program) Run(ctx context.Context) error {
	errCh := make(chan error, len(p.services))
	for _, svc := range p.services {
		svc := svc
		go func() { errCh <- svc.Run(ctx) }()
	}

	consoleDone := make(chan struct{})
	go func() {
		defer close(consoleDone)
		p.consoleLoop(ctx, os.Stdin, p.out)
	}()

	remaining := len(p.services)
	var lastErr error
	for remaining > 0 {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errCh:
			remaining--
			if err != nil && ctx.Err() == nil {
				lastErr = err
				slog.Error("service stopped", "err", err, "remaining", remaining)
			}
		}
	}
	if lastErr != nil {
		slog.Error("all services stopped; console still available", "last_err", lastErr)
	}

	select {
	case <-ctx.Done():
		return nil
	case <-consoleDone:
		// Keep the process alive unless the user explicitly exits the console.
		// Returning lastErr here would cause the whole server to exit ("flash quit")
		// when ports are occupied.
		return nil
	}
}

func (p *Program) Shutdown(ctx context.Context) {
	for _, svc := range p.services {
		_ = svc.Shutdown(ctx)
	}
}

func (p *Program) consoleLoop(ctx context.Context, in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	fmt.Fprintln(out, "console ready:")
	fmt.Fprintln(out, "  /account create <email> [reservedPlayerUid]")
	fmt.Fprintln(out, "  quit | exit")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if !sc.Scan() {
			return
		}
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" || line == "/exit" || line == "/quit" {
			slog.Info("console exit requested")
			return
		}
		err := p.commands.ExecuteLine(command.Context{Out: out, DataBase: p.db}, line)
		if err != nil {
			fmt.Fprintln(out, "error:", err.Error())
		}
	}
}
