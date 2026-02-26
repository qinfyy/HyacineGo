package config

import "time"

type Config struct {
	Admin      Admin      `json:"admin"`
	Dispatch   Dispatch   `json:"dispatch"`
	GameServer GameServer `json:"gameserver"`
	Logging    Logging    `json:"logging"`
	Debug      Debug      `json:"debug"`

	LoadedAt time.Time `json:"-"`
}

type Admin struct {
	Addr string `json:"addr"`
}

type Dispatch struct {
	Addr       string `json:"addr"`
	HotfixPath string `json:"hotfix_path"`
	FreesrDir  string `json:"freesr_dir"`
	CorsOrigin string `json:"cors_origin"`

	AccountPath       string `json:"account_path"`
	AutoCreateAccount bool   `json:"auto_create_account"`
}

type GameServer struct {
	Addr          string `json:"addr"`
	Protocol      string `json:"protocol"` // tcp|kcp
	MaxFrameBytes int    `json:"max_frame_bytes"`
	ResourceRoot  string `json:"resource_root"`
}

type Logging struct {
	Level  string `json:"level"`
	Format string `json:"format"` // json|text
	File   string `json:"file"`   // optional path
}

type Debug struct {
	AutoCreateAccount      bool     `json:"auto_create_account"`
	LogPackets             bool     `json:"log_packets"`
	LogUnknownPackets      bool     `json:"log_unknown_packets"`
	AutoReplyUnknown       *bool    `json:"auto_reply_unknown"`
	SendClientDownloadData bool     `json:"send_client_download_data"`
	LoginExtraSendCmdIds   []uint16 `json:"login_extra_send_cmd_ids"`
	TraceFile              string   `json:"trace_file"`
}
