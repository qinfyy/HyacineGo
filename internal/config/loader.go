package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Loader struct {
	path string
}

func NewLoader(path string) (*Loader, error) {
	if path == "" {
		return nil, fmt.Errorf("empty config path")
	}
	return &Loader{path: path}, nil
}

func defaultConfig() Config {
	autoReply := true
	return Config{
		Admin: Admin{
			Addr: "127.0.0.1:18080",
		},
		Dispatch: Dispatch{
			Addr:              "0.0.0.0:21000",
			HotfixPath:        "./hotfix.json",
			FreesrDir:         "./configs/freesr-data.json",
			CorsOrigin:        "*",
			AccountPath:       "./configs/accounts.json",
			AutoCreateAccount: true,
		},
		GameServer: GameServer{
			Addr:          "0.0.0.0:23301",
			Protocol:      "tcp",
			ResourceRoot:  "./resources",
			MaxFrameBytes: 1024 * 1024,
		},
		Logging: Logging{
			Level:  "debug",
			Format: "pretty",
			File:   "",
		},
		Debug: Debug{
			LogPackets:        true,
			LogUnknownPackets: true,
			AutoReplyUnknown:  &autoReply,
			AutoCreateAccount: true,
		},
		LoadedAt: time.Now(),
	}
}

func (l *Loader) Load() (Config, error) {
	raw, err := os.ReadFile(l.path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg := defaultConfig()
			data, _ := json.MarshalIndent(cfg, "", "  ")
			if err := os.WriteFile(l.path, data, 0644); err != nil {
				return cfg, fmt.Errorf("failed to create default config: %w", err)
			}
			return cfg, nil
		}
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal(raw, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.Admin.Addr == "" {
		cfg.Admin.Addr = "127.0.0.1:18080"
	}
	if cfg.Dispatch.Addr == "" {
		cfg.Dispatch.Addr = "0.0.0.0:21000"
	}
	if cfg.Dispatch.HotfixPath == "" {
		cfg.Dispatch.HotfixPath = "./hotfix.json"
	}
	if cfg.Dispatch.FreesrDir == "" {
		cfg.Dispatch.FreesrDir = "./configs/freesrdata"
	}
	if cfg.Dispatch.CorsOrigin == "" {
		cfg.Dispatch.CorsOrigin = "*"
	}
	if cfg.Dispatch.AccountPath == "" {
		cfg.Dispatch.AccountPath = "./configs/accounts.json"
	}
	if cfg.GameServer.Addr == "" {
		cfg.GameServer.Addr = "0.0.0.0:23301"
	}
	if cfg.GameServer.Protocol == "" {
		// Star Rail clients typically use KCP for the game connection.
		cfg.GameServer.Protocol = "kcp"
	}
	if cfg.GameServer.ResourceRoot == "" {
		cfg.GameServer.ResourceRoot = "./resources"
	}
	if cfg.GameServer.MaxFrameBytes <= 0 {
		cfg.GameServer.MaxFrameBytes = 1024 * 1024
	}
	if cfg.Logging.Level == "" {
		cfg.Logging.Level = "info"
	}
	if cfg.Logging.Format == "" {
		cfg.Logging.Format = "pretty"
	}
	// Developer-first defaults: keep debug on to avoid "silent missing packet" issues.
	if cfg.Logging.Level == "" || strings.EqualFold(cfg.Logging.Level, "info") {
		cfg.Logging.Level = "debug"
	}
	cfg.Debug.LogPackets = true
	cfg.Debug.LogUnknownPackets = true
	if cfg.Debug.AutoReplyUnknown == nil {
		v := true
		cfg.Debug.AutoReplyUnknown = &v
	}
	if !cfg.Dispatch.AutoCreateAccount && cfg.Debug.AutoCreateAccount {
		cfg.Dispatch.AutoCreateAccount = true
	}

	cfg.LoadedAt = time.Now()
	return cfg, nil
}
