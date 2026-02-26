package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	observability "hyacine-server/internal/common"
	"hyacine-server/internal/config"
	"hyacine-server/internal/dispatch"
	"hyacine-server/internal/gameserver/database"
	"hyacine-server/internal/gameserver/server"
	"hyacine-server/internal/gameserver/server/admin"
	"hyacine-server/internal/gameserver/util/program"
)

var (
	// buildTime is injected by the linker, e.g.:
	// go build -ldflags "-X main.buildTime=2025-12-23T21:03:00Z"
	buildTime = ""
	// buildCommit is injected by the linker, e.g.:
	// go build -ldflags "-X main.buildCommit=abcdef0"
	buildCommit = ""
)

func buildInfo() (timeStr, commit string) {
	timeStr = buildTime
	commit = buildCommit
	if bi, ok := debug.ReadBuildInfo(); ok && bi != nil {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.time":
				if timeStr == "" {
					timeStr = s.Value
				}
			case "vcs.revision":
				if commit == "" {
					commit = s.Value
				}
			}
		}
	}
	if timeStr == "" {
		timeStr = "unknown"
	}
	if commit == "" {
		commit = "unknown"
	}
	return timeStr, commit
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./configs/config.json", "path to config json")
	flag.Parse()

	logger, closeLog, err := observability.ConfigureLogger(observability.LoggerOptions{
		Level:  "info",
		Format: "json",
	})
	if err != nil {
		slog.Error("logger init failed", "err", err)
		os.Exit(1)
	}
	defer func() {
		if closeLog != nil {
			closeLog()
		}
	}()
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	loader, err := config.NewLoader(configPath)
	if err != nil {
		slog.Error("config loader init failed", "err", err)
		os.Exit(1)
	}

	cfg, err := loader.Load()
	if err != nil {
		slog.Error("config load failed", "err", err, "path", configPath)
		os.Exit(1)
	}
	if l, c, err := observability.ConfigureLogger(observability.LoggerOptions{
		Level:  cfg.Logging.Level,
		Format: cfg.Logging.Format,
		File:   cfg.Logging.File,
	}); err == nil {
		slog.SetDefault(l)
		if c != nil {
			defer c()
		}
	} else {
		observability.SetLevel(cfg.Logging.Level)
	}

	metrics := observability.NewMetrics()

	bt, bc := buildInfo()
	slog.Info("build", "time", bt, "commit", bc)

	db := database.New(database.Options{
		AccountPath: cfg.Dispatch.AccountPath,
		PlayerDir:   "./configs/players",
	})
	if err := db.Load(); err != nil {
		slog.Warn("database load failed", "err", err)
	}

	dispatchAutoCreate := cfg.Dispatch.AutoCreateAccount || cfg.Debug.AutoCreateAccount
	dispatchServer := dispatch.New(dispatch.Options{
		Addr:              cfg.Dispatch.Addr,
		GameAddr:          cfg.GameServer.Addr,
		HotfixPath:        cfg.Dispatch.HotfixPath,
		FreesrDir:         cfg.Dispatch.FreesrDir,
		CorsOrigin:        cfg.Dispatch.CorsOrigin,
		AccountPath:       cfg.Dispatch.AccountPath,
		AutoCreateAccount: dispatchAutoCreate,
		Metrics:           metrics,
		DataBase:          db,
	}, metrics)

	autoReplyUnknown := true
	if cfg.Debug.AutoReplyUnknown != nil {
		autoReplyUnknown = *cfg.Debug.AutoReplyUnknown
	}
	gameServer := server.New(cfg.GameServer.Addr, server.Options{
		Protocol:               cfg.GameServer.Protocol,
		MaxFrameBytes:          cfg.GameServer.MaxFrameBytes,
		AccountPath:            cfg.Dispatch.AccountPath,
		DataBase:               db,
		LogPackets:             cfg.Debug.LogPackets,
		LogUnknown:             cfg.Debug.LogUnknownPackets,
		AutoReply:              autoReplyUnknown,
		SendClientDownloadData: cfg.Debug.SendClientDownloadData,
		LoginExtraSendCmdIDs:   cfg.Debug.LoginExtraSendCmdIds,
		TraceFile:              cfg.Debug.TraceFile,
		ResourceRoot:           cfg.GameServer.ResourceRoot,
		FreesrDir:              cfg.Dispatch.FreesrDir,
	}, metrics)

	// Hot reload used by both admin server and dispatch /admin/reload.
	var reloadFn func() error
	reloadFn = func() error {
		newCfg, err := loader.Load()
		if err != nil {
			return err
		}
		observability.SetLevel(newCfg.Logging.Level)
		if newCfg.Dispatch.AccountPath != cfg.Dispatch.AccountPath {
			db = database.New(database.Options{AccountPath: newCfg.Dispatch.AccountPath, PlayerDir: "./configs/players"})
		}
		_ = db.Load()
		cfg = newCfg
		dispatchServer.UpdateOptions(dispatch.Options{
			Addr:              newCfg.Dispatch.Addr,
			GameAddr:          newCfg.GameServer.Addr,
			HotfixPath:        newCfg.Dispatch.HotfixPath,
			FreesrDir:         newCfg.Dispatch.FreesrDir,
			CorsOrigin:        newCfg.Dispatch.CorsOrigin,
			AccountPath:       newCfg.Dispatch.AccountPath,
			AutoCreateAccount: newCfg.Dispatch.AutoCreateAccount || newCfg.Debug.AutoCreateAccount,
			Reload:            reloadFn,
			Metrics:           metrics,
			DataBase:          db,
		})
		autoReplyUnknown := true
		if newCfg.Debug.AutoReplyUnknown != nil {
			autoReplyUnknown = *newCfg.Debug.AutoReplyUnknown
		}
		gameServer.UpdateOptions(server.Options{
			Protocol:               newCfg.GameServer.Protocol,
			MaxFrameBytes:          newCfg.GameServer.MaxFrameBytes,
			AccountPath:            newCfg.Dispatch.AccountPath,
			DataBase:               db,
			LogPackets:             newCfg.Debug.LogPackets,
			LogUnknown:             newCfg.Debug.LogUnknownPackets,
			AutoReply:              autoReplyUnknown,
			SendClientDownloadData: newCfg.Debug.SendClientDownloadData,
			LoginExtraSendCmdIDs:   newCfg.Debug.LoginExtraSendCmdIds,
			TraceFile:              newCfg.Debug.TraceFile,
			ResourceRoot:           newCfg.GameServer.ResourceRoot,
			FreesrDir:              newCfg.Dispatch.FreesrDir,
		})
		return nil
	}
	dispatchServer.UpdateOptions(dispatch.Options{
		Addr:              cfg.Dispatch.Addr,
		GameAddr:          cfg.GameServer.Addr,
		HotfixPath:        cfg.Dispatch.HotfixPath,
		FreesrDir:         cfg.Dispatch.FreesrDir,
		CorsOrigin:        cfg.Dispatch.CorsOrigin,
		AccountPath:       cfg.Dispatch.AccountPath,
		AutoCreateAccount: dispatchAutoCreate,
		Reload:            reloadFn,
		Metrics:           metrics,
		DataBase:          db,
	})

	adminServer := admin.New(admin.Options{
		Addr:       cfg.Admin.Addr,
		ConfigPath: configPath,
		Reload:     reloadFn,
		Metrics:    metrics,
	})

	prg := program.New(program.Options{
		Services: []program.Service{dispatchServer, gameServer, adminServer},
		DataBase: db,
		Out:      os.Stdout,
	})
	errCh := make(chan error, 1)
	go func() { errCh <- prg.Run(ctx) }()

	slog.Info("server started",
		"admin", cfg.Admin.Addr,
		"dispatch", cfg.Dispatch.Addr,
		"gameserver", cfg.GameServer.Addr,
	)

	select {
	case <-ctx.Done():
		slog.Info("shutdown requested")
	case err := <-errCh:
		slog.Error("server stopped with error", "err", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	prg.Shutdown(shutdownCtx)

	fmt.Fprintln(os.Stdout, "bye")
}
