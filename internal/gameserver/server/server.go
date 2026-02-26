package server

import (
	"context"
	"encoding/binary"
	"errors"
	"log/slog"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	observability "hyacine-server/internal/common"
	"hyacine-server/internal/gameserver/data"
	"hyacine-server/internal/gameserver/database"
	"hyacine-server/internal/gameserver/freesrdata"
	"hyacine-server/internal/gameserver/game/scene"
	"hyacine-server/internal/gameserver/server/packet"
	"hyacine-server/internal/gameserver/server/packet/recv"
	"hyacine-server/internal/gameserver/server/trace"

	"google.golang.org/protobuf/proto"
)

type Options struct {
	Protocol               string
	MaxFrameBytes          int
	AccountPath            string
	DataBase               *database.DataBase
	LogPackets             bool
	LogUnknown             bool
	AutoReply              bool
	SendClientDownloadData bool
	LoginExtraSendCmdIDs   []uint16
	TraceFile              string
	ResourceRoot           string
	FreesrDir              string
}

type Server struct {
	addr string

	mu      sync.RWMutex
	opts    Options
	metrics *observability.Metrics

	ln  net.Listener
	udp *net.UDPConn

	reg *packet.Registry
	db  *database.DataBase

	tracer       *trace.Tracer
	data         *data.GameData
	freesr       *freesrdata.Data
	downloadData []byte
	scene        scene.DefaultScene
}

func (s *Server) registerHandlers() {
	recv.RegisterAll(s)
}

func New(addr string, opts Options, metrics *observability.Metrics) *Server {
	if opts.MaxFrameBytes <= 0 {
		opts.MaxFrameBytes = 1024 * 1024
	}
	if opts.AccountPath == "" {
		opts.AccountPath = "./configs/accounts.json"
	}
	if strings.TrimSpace(opts.Protocol) == "" {
		opts.Protocol = "kcp"
	}

	s := &Server{addr: addr, opts: opts, metrics: metrics}
	if opts.DataBase != nil {
		s.db = opts.DataBase
	} else {
		s.db = database.New(database.Options{AccountPath: opts.AccountPath, PlayerDir: "./configs/players"})
	}
	if err := s.db.Load(); err != nil {
		slog.Warn("gameserver database load failed", "err", err)
	}
	s.reg = packet.NewRegistry()
	s.registerHandlers()
	if strings.TrimSpace(opts.TraceFile) != "" {
		if tr, err := trace.New(opts.TraceFile); err == nil {
			s.tracer = tr
		} else {
			slog.Warn("trace init failed", "err", err, "path", opts.TraceFile)
		}
	}
	resourceRoot := strings.TrimSpace(opts.ResourceRoot)
	if resourceRoot == "" {
		resourceRoot = "./resources"
	}
	s.data = data.NewGameData(data.NewResourceLoader(resourceRoot))

	if strings.TrimSpace(opts.FreesrDir) != "" {
		if fs, err := freesrdata.LoadDir(opts.FreesrDir); err == nil {
			s.freesr = fs
			slog.Debug("freesrdata loaded", "dir", opts.FreesrDir, "avatars", fs.AvatarCount())
		} else {
			slog.Warn("freesrdata load failed", "dir", opts.FreesrDir, "err", err)
		}
	}
	if sc, err := scene.PickDefault(s.data); err == nil {
		s.scene = sc
	} else {
		// 使用与PickDefault相同的默认值
		s.scene = scene.DefaultScene{EntryID: 2050201, PlaneID: 20502, FloorID: 20502001, WorldID: 401}
		slog.Warn("使用默认场景配置，因为PickDefault失败", "err", err)
	}
	s.warmupGameData()
	return s
}

func (s *Server) warmupGameData() {
	if s.data == nil {
		return
	}

	if t, err := s.data.AvatarConfig(); err != nil {
		slog.Warn("data load failed", "table", "AvatarConfig", "err", err)
	} else {
		slog.Debug("data loaded", "table", "AvatarConfig", "rows", len(t))
	}
	if t, err := s.data.MapEntrance(); err != nil {
		slog.Warn("data load failed", "table", "MapEntrance", "err", err)
	} else {
		slog.Debug("data loaded", "table", "MapEntrance", "rows", len(t))
	}
	if t, err := s.data.MazePlane(); err != nil {
		slog.Warn("data load failed", "table", "MazePlane", "err", err)
	} else {
		slog.Debug("data loaded", "table", "MazePlane", "rows", len(t))
	}
	if t, err := s.data.MappingInfo(); err != nil {
		slog.Warn("data load failed", "table", "MappingInfo", "err", err)
	} else {
		slog.Debug("data loaded", "table", "MappingInfo", "rows", len(t))
	}
	if t, err := s.data.NpcMonsterData(); err != nil {
		slog.Warn("data load failed", "table", "NPCMonsterData", "err", err)
	} else {
		slog.Debug("data loaded", "table", "NPCMonsterData", "rows", len(t))
	}
	if t, err := s.data.MonsterConfig(); err != nil {
		slog.Warn("data load failed", "table", "MonsterConfig", "err", err)
	} else {
		slog.Debug("data loaded", "table", "MonsterConfig", "rows", len(t))
	}
	if t, err := s.data.TeleportConfig(); err != nil {
		slog.Warn("data load failed", "table", "TeleportConfig", "err", err)
	} else {
		slog.Debug("data loaded", "table", "TeleportConfig", "rows", len(t))
	}

	// 预加载和验证默认场景的楼层信息
	if s.scene.FloorID != 0 {
		if floor, err := s.data.LevelOutputFloor(int(s.scene.FloorID)); err != nil {
			slog.Warn("默认场景楼层加载失败", "floorID", s.scene.FloorID, "err", err)
		} else if floor == nil {
			slog.Warn("默认场景楼层为空", "floorID", s.scene.FloorID)
		} else {
			slog.Debug("默认场景楼层加载成功", "floorID", s.scene.FloorID, "groups", len(floor.GroupInstanceList))

			// 验证组路径
			for _, gi := range floor.GroupInstanceList {
				if gi.IsDelete || gi.GroupPath == "" || gi.ID == 0 {
					continue
				}
				if _, err := s.data.LoadLevelOutputGroupByPath(gi.GroupPath); err != nil {
					slog.Warn("组加载失败", "groupID", gi.ID, "groupPath", gi.GroupPath, "err", err)
				}
			}
		}
	}
}

func (s *Server) UpdateOptions(opts Options) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if strings.TrimSpace(opts.Protocol) != "" && !strings.EqualFold(opts.Protocol, s.opts.Protocol) {
		s.opts.Protocol = opts.Protocol
		slog.Warn("gameserver protocol updated (restart required to take effect)", "protocol", opts.Protocol)
	}
	if opts.MaxFrameBytes > 0 {
		s.opts.MaxFrameBytes = opts.MaxFrameBytes
	}
	if opts.FreesrDir != "" && opts.FreesrDir != s.opts.FreesrDir {
		s.opts.FreesrDir = opts.FreesrDir
		if fs, err := freesrdata.LoadDir(opts.FreesrDir); err == nil {
			s.freesr = fs
			slog.Debug("freesrdata reloaded", "dir", opts.FreesrDir, "avatars", fs.AvatarCount())
		} else {
			slog.Warn("freesrdata reload failed", "dir", opts.FreesrDir, "err", err)
		}
	}
	if opts.DataBase != nil {
		s.db = opts.DataBase
		s.opts.DataBase = opts.DataBase
		_ = s.db.Load()
	}
	if opts.AccountPath != "" && opts.AccountPath != s.opts.AccountPath {
		s.opts.AccountPath = opts.AccountPath
		if s.db != nil {
			_ = s.db.Load()
		}
	}
}

func (s *Server) getClientDownloadData() []byte {
	s.mu.RLock()
	cached := s.downloadData
	s.mu.RUnlock()
	if cached != nil {
		return cached
	}

	// Mirrors 3.8.0src behavior: send a Lua payload via ClientDownloadDataScNotify.
	// Prefer a file so users can easily swap scripts without recompiling.
	paths := []string{
		filepath.Clean("./configs/client_download.lua"),
		filepath.Clean("./configs/Hueby.json"), // 3.8.0src heartbeat example
		filepath.Clean("./configs/Hueby.lua"),
	}
	var data []byte
	for _, p := range paths {
		if b, err := os.ReadFile(p); err == nil && len(b) > 0 {
			data = b
			break
		}
	}
	if data == nil {
		data = []byte("return\n")
	}

	s.mu.Lock()
	if s.downloadData == nil {
		s.downloadData = data
	}
	cached = s.downloadData
	s.mu.Unlock()
	return cached
}

func (s *Server) Run(ctx context.Context) error {
	s.mu.RLock()
	proto := strings.ToLower(strings.TrimSpace(s.opts.Protocol))
	s.mu.RUnlock()
	if proto == "" {
		proto = "kcp"
	}

	if proto == "kcp" {
		return s.runKCP(ctx)
	}

	var ln net.Listener
	var err error
	switch proto {
	case "tcp":
		ln, err = net.Listen("tcp", s.addr)
	default:
		return errors.New("unsupported gameserver protocol: " + proto)
	}
	if err != nil {
		return err
	}
	s.ln = ln

	slog.Info("gameserver listening", "addr", s.addr, "net", proto)

	go func() {
		<-ctx.Done()
		_ = s.Shutdown(context.Background())
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		if s.metrics != nil {
			s.metrics.IncGameConnections()
		}
		slog.Info("gameserver accepted", "remote", conn.RemoteAddr().String(), "local", conn.LocalAddr().String(), "net", proto)
		go s.handleConn(ctx, conn)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	_ = ctx
	if s.tracer != nil {
		_ = s.tracer.Close()
	}
	if s.udp != nil {
		_ = s.udp.Close()
	}
	if s.ln == nil {
		return nil
	}
	return s.ln.Close()
}

type session struct {
	conn net.Conn
	ps   packet.Session
}

func (s *Server) handleConn(ctx context.Context, c net.Conn) {
	start := time.Now()
	defer func() {
		slog.Info("gameserver disconnected", "remote", c.RemoteAddr().String(), "dur_ms", time.Since(start).Milliseconds())
		_ = c.Close()
	}()

	sess := &session{
		conn: c,
		ps: packet.Session{
			State:      packet.StateWaitingForToken,
			LastActive: time.Now(),
		},
	}
	readBuf := make([]byte, 64*1024)
	var pending []byte

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		_ = c.SetReadDeadline(time.Now().Add(120 * time.Second))
		n, err := c.Read(readBuf)
		if err != nil {
			return
		}
		if n <= 0 {
			continue
		}
		pending = append(pending, readBuf[:n]...)

		for {
			cmdID, payload, rest, ok := tryDecodePacket(pending)
			if !ok {
				break
			}
			pending = rest
			s.handlePacket(sess, cmdID, payload)
		}
	}
}

func (s *Server) handlePacket(sess *session, cmdID uint16, payload []byte) {
	if s.metrics != nil {
		s.metrics.IncGameFrames()
	}

	s.mu.RLock()
	logPackets := s.opts.LogPackets
	logUnknown := s.opts.LogUnknown
	s.mu.RUnlock()
	ps := &sess.ps
	ps.LastActive = time.Now()

	if logPackets {
		slog.Info("RECV", "cmd", packet.Name(cmdID), "cmd_id", cmdID, "bytes", len(payload), "state", ps.State.String())
	}
	if s.tracer != nil {
		s.tracer.Write(trace.Event{Dir: "RECV", CmdID: cmdID, Cmd: packet.Name(cmdID), Known: packet.Known(cmdID), Bytes: len(payload), State: ps.State.String()})
	}
	ctx := packet.Context{
		Session: ps,
		Send: func(outCmd uint16, msg proto.Message) {
			s.sendPacket(sess, outCmd, msg)
		},
		SendEmpty: func(outCmd uint16) {
			s.sendEmptyPacket(sess, outCmd)
		},
		Now: time.Now,
		Log: slog.Default(),
	}

	handled := s.reg.Handle(ctx, cmdID, payload)
	if handled {
		return
	}

	if logUnknown {
		slog.Warn("UNHANDLED", "cmd", packet.Name(cmdID), "cmd_id", cmdID, "known", packet.Known(cmdID), "bytes", len(payload), "state", ps.State.String())
	}
	if s.tracer != nil {
		s.tracer.Write(trace.Event{Dir: "UNHANDLED", CmdID: cmdID, Cmd: packet.Name(cmdID), Known: packet.Known(cmdID), Bytes: len(payload), State: ps.State.String()})
	}

	s.mu.RLock()
	autoReply := s.opts.AutoReply
	s.mu.RUnlock()
	if autoReply {
		// AutoReply is a compatibility fallback: reply with an ScRsp when a handler is missing.
		// Do NOT echo an empty packet for the incoming CsReq cmd id: that produces confusing "SEND CmdId" noise
		// and can break some clients that expect only ScRsp/notify packets from server.
		if rspID, m, ok := packet.AutoReplyScRsp(cmdID); ok {
			s.sendPacket(sess, rspID, m)
			if s.tracer != nil {
				s.tracer.Write(trace.Event{Dir: "AUTOREPLY", CmdID: rspID, Cmd: packet.Name(rspID), Known: packet.Known(rspID), Bytes: 0, State: ps.State.String(), Note: "ScRsp"})
			}
		} else if rspID, ok := packet.ScRspForCsReq(cmdID); ok {
			s.sendEmptyPacket(sess, rspID)
			if s.tracer != nil {
				s.tracer.Write(trace.Event{Dir: "AUTOREPLY", CmdID: rspID, Cmd: packet.Name(rspID), Known: packet.Known(rspID), Bytes: 0, State: ps.State.String(), Note: "ScRsp(empty)"})
			}
		}
	}
}

func (s *Server) sendPacket(sess *session, cmdID uint16, m proto.Message) {
	data, err := proto.Marshal(m)
	if err != nil {
		return
	}
	packetBytes := buildPacket(cmdID, data)

	s.mu.RLock()
	logPackets := s.opts.LogPackets
	s.mu.RUnlock()
	if logPackets {
		slog.Info("SEND", "cmd", packet.Name(cmdID), "cmd_id", cmdID, "bytes", len(data))
	}
	if s.tracer != nil {
		s.tracer.Write(trace.Event{Dir: "SEND", CmdID: cmdID, Cmd: packet.Name(cmdID), Known: packet.Known(cmdID), Bytes: len(data)})
	}

	_ = sess.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, _ = sess.conn.Write(packetBytes)
}

func (s *Server) sendEmptyPacket(sess *session, cmdID uint16) {
	packetBytes := buildPacket(cmdID, nil)
	s.mu.RLock()
	logPackets := s.opts.LogPackets
	s.mu.RUnlock()
	if logPackets {
		slog.Info("SEND", "cmd", packet.Name(cmdID), "cmd_id", cmdID, "bytes", 0)
	}
	if s.tracer != nil {
		s.tracer.Write(trace.Event{Dir: "SEND", CmdID: cmdID, Cmd: packet.Name(cmdID), Known: packet.Known(cmdID), Bytes: 0})
	}
	_ = sess.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, _ = sess.conn.Write(packetBytes)
}

const (
	headerConst = 0x9d74c714
	tailConst   = 0xd7a152c8
)

func tryDecodePacket(b []byte) (cmdID uint16, payload []byte, rest []byte, ok bool) {
	if len(b) < 16 {
		return 0, nil, b, false
	}
	if binary.BigEndian.Uint32(b[0:4]) != headerConst {
		return 0, nil, nil, false
	}
	cmdID = binary.BigEndian.Uint16(b[4:6])
	headerLen := int(binary.BigEndian.Uint16(b[6:8]))
	dataLen := int(binary.BigEndian.Uint32(b[8:12]))
	total := 4 + 2 + 2 + 4 + headerLen + dataLen + 4
	if total < 16 || total > len(b) {
		return 0, nil, b, false
	}
	tailPos := total - 4
	if binary.BigEndian.Uint32(b[tailPos:total]) != tailConst {
		return 0, nil, nil, false
	}
	payloadStart := 12 + headerLen
	payloadEnd := payloadStart + dataLen
	return cmdID, b[payloadStart:payloadEnd], b[total:], true
}

func buildPacket(cmdID uint16, payload []byte) []byte {
	headerLen := uint16(0)
	dataLen := uint32(len(payload))
	out := make([]byte, 16+len(payload))
	binary.BigEndian.PutUint32(out[0:4], headerConst)
	binary.BigEndian.PutUint16(out[4:6], cmdID)
	binary.BigEndian.PutUint16(out[6:8], headerLen)
	binary.BigEndian.PutUint32(out[8:12], dataLen)
	if len(payload) > 0 {
		copy(out[12:], payload)
	}
	binary.BigEndian.PutUint32(out[12+len(payload):], tailConst)
	return out
}
