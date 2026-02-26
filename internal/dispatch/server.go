package dispatch

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	observability "hyacine-server/internal/common"
	"hyacine-server/internal/gameserver/database"
	"hyacine-server/internal/gameserver/util/srtools"
	pb "hyacine-server/internal/proto/gen"

	"google.golang.org/protobuf/proto"
)

type Options struct {
	Addr       string
	GameAddr   string
	HotfixPath string
	FreesrDir  string
	CorsOrigin string

	AccountPath       string
	AutoCreateAccount bool
	DataBase          *database.DataBase

	Reload  func() error
	Metrics *observability.Metrics
}

type Server struct {
	mu      sync.RWMutex
	opts    Options
	metrics *observability.Metrics
	db      *database.DataBase
	srtools *srtools.Handler

	srv *http.Server
	ln  net.Listener
}

func New(opts Options, metrics *observability.Metrics) *Server {
	s := &Server{
		opts:    withDispatchDefaults(opts),
		metrics: metrics,
	}
	if s.opts.DataBase != nil {
		s.db = s.opts.DataBase
	} else {
		s.db = database.New(database.Options{AccountPath: s.opts.AccountPath})
	}
	if err := s.db.Load(); err != nil {
		slog.Warn("database load failed", "err", err)
	}
	s.srtools = srtools.New(srtools.Options{
		Dir:        s.opts.FreesrDir,
		CorsOrigin: s.opts.CorsOrigin,
	})
	s.rebuildHTTP()
	return s
}

func (s *Server) UpdateOptions(opts Options) {
	s.mu.Lock()
	defer s.mu.Unlock()

	opts = withDispatchDefaults(opts)
	changedAddr := opts.Addr != "" && opts.Addr != s.opts.Addr
	s.opts = opts
	if s.opts.DataBase != nil {
		s.db = s.opts.DataBase
	} else if s.db == nil {
		s.db = database.New(database.Options{AccountPath: s.opts.AccountPath})
	}
	if s.db != nil {
		_ = s.db.Load()
	}
	if s.srtools != nil {
		s.srtools = srtools.New(srtools.Options{
			Dir:        s.opts.FreesrDir,
			CorsOrigin: s.opts.CorsOrigin,
		})
	}
	if changedAddr {
		s.rebuildHTTP()
		slog.Warn("dispatch addr updated (restart required to take effect)", "addr", opts.Addr)
	}
}

func withDispatchDefaults(opts Options) Options {
	if opts.Addr == "" {
		opts.Addr = "0.0.0.0:21000"
	}
	if opts.GameAddr == "" {
		opts.GameAddr = "0.0.0.0:22102"
	}
	if opts.CorsOrigin == "" {
		opts.CorsOrigin = "*"
	}
	if opts.AccountPath == "" {
		opts.AccountPath = "./configs/accounts.json"
	}
	return opts
}

func (s *Server) rebuildHTTP() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.route)

	s.srv = &http.Server{
		Addr:              s.opts.Addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.mu.RLock()
	srv := s.srv
	s.mu.RUnlock()

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	go func() {
		<-ctx.Done()
		_ = s.Shutdown(context.Background())
	}()

	err = srv.Serve(ln)
	if err == nil || errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.srv == nil {
		return nil
	}
	return s.srv.Shutdown(ctx)
}

func (s *Server) route(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ww := &statusWriter{ResponseWriter: w, status: http.StatusOK}
	defer func() {
		slog.Info("http", "component", "dispatch", "method", r.Method, "path", r.URL.Path, "status", ww.status, "dur_ms", time.Since(start).Milliseconds(), "remote", r.RemoteAddr)
	}()

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/query_dispatch":
		s.handleQueryDispatch(ww, r)
	case r.Method == http.MethodGet && r.URL.Path == "/query_gateway":
		s.handleQueryGateway(ww, r)
	case r.Method == http.MethodPost && r.URL.Path == "/account/risky/api/check":
		s.handleRiskyCheck(ww, r)
	case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/mdk/shield/api/login"):
		s.handleShieldLogin(ww, r)
	case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/mdk/shield/api/verify"):
		s.handleShieldVerify(ww, r)
	case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/mdk/shield/api/getConfig"):
		s.handleShieldGetConfig(ww, r)
	case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/mdk/shield/api/loadConfig"):
		s.handleShieldLoadConfig(ww, r)
	case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/mdk/shield/api/loginByPassword"):
		s.handleLoginByPassword(ww, r)
	case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/combo/granter/login/v2/login"):
		s.handleComboLogin(ww, r)
	case r.Method == http.MethodOptions && r.URL.Path == "/srtools":
		s.srtools.Options(ww, r)
	case r.Method == http.MethodPost && r.URL.Path == "/srtools":
		s.srtools.Save(ww, r)
	case r.Method == http.MethodGet && r.URL.Path == "/admin/healthz":
		s.handleAdminHealthz(ww, r)
	case r.Method == http.MethodPost && r.URL.Path == "/admin/reload":
		s.handleAdminReload(ww, r)
	case r.Method == http.MethodGet && r.URL.Path == "/admin/metrics":
		s.handleAdminMetrics(ww, r)
	default:
		http.NotFound(ww, r)
	}
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (s *Server) handleQueryDispatch(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()

	dispatchURL := s.dispatchURL(r)
	msg := &pb.Dispatch{
		Retcode: 0,
		RegionList: []*pb.RegionInfo{
			{
				Name:        "HyacineGo",
				Title:       "HyacineGo",
				EnvType:     "9",
				DisplayName: "HyacineGo",
				DispatchUrl: dispatchURL,
			},
		},
	}
	writeB64Proto(w, msg)
}

func (s *Server) handleQueryGateway(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()

	s.mu.RLock()
	gameAddr := s.opts.GameAddr
	hotfixPath := s.opts.HotfixPath
	s.mu.RUnlock()

	ip, port, err := parseHostPort(gameAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	version := r.URL.Query().Get("version")
	hotfix := loadHotfixForVersion(hotfixPath, version)

	msg := &pb.GateServer{
		RegionName:      "local",
		Ip:              ip,
		Port:            uint32(port),
		AssetBundleUrl:  hotfix.AssetBundleURL,
		ExResourceUrl:   hotfix.ExResourceURL,
		LuaUrl:          hotfix.LuaURL,
		IfixUrl:         hotfix.IfixURL,
		IfixVersion:     hotfix.IfixVersion,
		MdkResVersion:   hotfix.MdkResVersion,
		ClientSecretKey: hotfix.ClientSecretKey,
		Msg:             hotfix.Msg,
		Unk1:            true,
		Unk2:            true,
		Unk3:            true,
		Unk4:            true,
		Unk5:            true,
		Unk6:            true,
		Unk7:            true,
		Unk8:            true,
		Unk9:            true,
		Unk10:           true,
		Unk11:           true,
		Unk12:           true,
		Unk13:           true,
		Unk14:           true,
		Unk15:           true,
	}
	writeB64Proto(w, msg)
}

func (s *Server) handleAdminHealthz(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func (s *Server) handleAdminReload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	s.mu.RLock()
	reload := s.opts.Reload
	s.mu.RUnlock()
	if reload == nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	if err := reload(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("reloaded\n"))
}

func (s *Server) handleAdminMetrics(w http.ResponseWriter, r *http.Request) {
	_ = r
	s.mu.RLock()
	m := s.opts.Metrics
	s.mu.RUnlock()
	if m == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	m.Handler().ServeHTTP(w, r)
}

type hotfixEntry struct {
	AssetBundleURL  string `json:"asset_bundle_url"`
	ExResourceURL   string `json:"ex_resource_url"`
	LuaURL          string `json:"lua_url"`
	IfixURL         string `json:"ifix_url"`
	IfixVersion     string `json:"ifix_version"`
	MdkResVersion   string `json:"mdk_res_version"`
	ClientSecretKey string `json:"client_secret_key"`
	Msg             string `json:"msg"`
}

func loadHotfixForVersion(path string, version string) hotfixEntry {
	if strings.TrimSpace(version) == "" {
		return hotfixEntry{}
	}

	// When links are empty (or hotfix isn't configured), reply with an empty hotfix payload.
	// This avoids the official hotfix flow (slow + easier to detect).
	path = strings.TrimSpace(path)
	if path == "" {
		return hotfixEntry{}
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return hotfixEntry{}
	}

	var root map[string]hotfixEntry
	if err := json.Unmarshal(raw, &root); err != nil {
		return hotfixEntry{}
	}

	// Exact match.
	if v, ok := root[version]; ok {
		return v
	}
	// Optional wildcard/default.
	if v, ok := root["*"]; ok {
		return v
	}
	return hotfixEntry{}
}

func (s *Server) handleRiskyCheck(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"id":      "06611ed14c3131a676b19c0d34c0644b",
			"action":  "ACTION_NONE",
			"geetest": nil,
		},
	})
}

func (s *Server) handleShieldLogin(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	a := s.ensureAccount(r)
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"account": map[string]any{
				"area_code":       "**",
				"email":           a.Email,
				"country":         "CN",
				"is_email_verify": "1",
				"token":           a.DispatchToken,
				"uid":             a.UID,
			},
			"device_grant_required": false,
			"reactivate_required":   false,
			"realperson_required":   false,
			"safe_mobilerequired":   false,
		},
	})
}

func (s *Server) handleShieldVerify(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	a := s.ensureAccount(r)
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"account": map[string]any{
				"area_code":       "**",
				"email":           a.Email,
				"country":         "CN",
				"is_email_verify": "1",
				"token":           a.DispatchToken,
				"uid":             a.UID,
			},
			"device_grant_required": false,
			"reactivate_required":   false,
			"realperson_required":   false,
			"safe_mobilerequired":   false,
		},
	})
}

func (s *Server) handleComboLogin(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	a := s.ensureAccount(r)
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"account_type": 1,
			"open_id":      a.UID,
			"combo_id":     a.UID,
			"combo_token":  a.ComboToken,
			"heartbeat":    false,
			"data":         "{\"guest\": false}",
		},
	})
}

func (s *Server) handleShieldGetConfig(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	_ = r
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"protocol":                  true,
			"qr_enabled":                false,
			"log_level":                 "INFO",
			"announce_url":              "",
			"push_alias_type":           0,
			"disable_ysdk_guard":        true,
			"enable_announce_pic_popup": false,
			"app_name":                  "hkrpg_global",
			"qr_enabled_apps":           map[string]any{"bbs": false, "cloud": false},
			"qr_app_icons":              map[string]any{"app": "", "bbs": "", "cloud": ""},
			"qr_cloud_display_name":     "",
			"enable_user_center":        true,
			"functional_switch_configs": map[string]any{},
		},
	})
}

func (s *Server) handleShieldLoadConfig(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	_ = r
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"id":                     24,
			"game_key":               "hkrpg_global",
			"client":                 "PC",
			"identity":               "I_IDENTITY",
			"guest":                  false,
			"ignore_versions":        "",
			"scene":                  "S_NORMAL",
			"name":                   "hkrpg_global",
			"disable_regist":         false,
			"enable_email_captcha":   false,
			"thirdparty":             []string{"fb", "tw", "gl", "ap"},
			"disable_mmt":            false,
			"server_guest":           false,
			"thirdparty_ignore":      []string{},
			"enable_ps_bind_account": false,
			"thirdparty_login_configs": map[string]any{
				"tw": map[string]any{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 2592000},
				"ap": map[string]any{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 604800},
				"fb": map[string]any{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 2592000},
				"gl": map[string]any{"token_type": "TK_GAME_TOKEN", "game_token_expires_in": 604800},
			},
			"initialize_firebase":   false,
			"bbs_auth_login":        false,
			"bbs_auth_login_ignore": map[string]any{},
			"fetch_instance_id":     false,
			"enable_flash_login":    false,
		},
	})
}

func (s *Server) handleLoginByPassword(w http.ResponseWriter, r *http.Request) {
	s.incDispatch()
	a := s.ensureAccount(r)
	writeJSON(w, http.StatusOK, map[string]any{
		"retcode": 0,
		"message": "OK",
		"data": map[string]any{
			"bind_email_action_ticket": "",
			"ext_user_info":            map[string]any{"birth": "0", "guardian_email": ""},
			"reactivate_action_token":  "",
			"token":                    map[string]any{"token": a.DispatchToken, "token_type": "1"},
			"user_info": map[string]any{
				"account_name":    "HyacineGo",
				"aid":             a.UID,
				"area_code":       "**",
				"country":         "CN",
				"email":           a.Email,
				"is_email_verify": "1",
			},
		},
	})
}

type loginAccount struct {
	UID           string
	Email         string
	DispatchToken string
	ComboToken    string
}

func (s *Server) ensureAccount(r *http.Request) loginAccount {
	s.mu.RLock()
	auto := s.opts.AutoCreateAccount
	s.mu.RUnlock()

	username := extractUsername(r)
	if username == "" {
		username = "HyacineLover@StarRail.com"
	}

	var aUID, aEmail, aDispatchToken, aComboToken string
	if s.db != nil {
		if a, ok := s.db.GetAccountByUsername(username); ok && a != nil {
			aUID, aEmail, aDispatchToken, aComboToken = a.UID, a.Email(), a.DispatchToken, a.ComboToken
		} else if auto {
			a, _, err := s.db.GetOrCreateAccount(username)
			if err == nil && a != nil {
				_ = s.db.SaveAccounts()
				aUID, aEmail, aDispatchToken, aComboToken = a.UID, a.Email(), a.DispatchToken, a.ComboToken
			}
		}
	}
	if aUID == "" {
		// fallback: non-persistent ephemeral identity
		return loginAccount{
			UID:           "1001",
			Email:         username,
			DispatchToken: "aa",
			ComboToken:    "bb",
		}
	}
	return loginAccount{
		UID:           aUID,
		Email:         aEmail,
		DispatchToken: aDispatchToken,
		ComboToken:    aComboToken,
	}
}

func extractUsername(r *http.Request) string {
	// Best-effort extraction; real servers validate and support more flows.
	var body struct {
		Account string `json:"account"`
		Email   string `json:"email"`
		UID     string `json:"uid"`
		Token   string `json:"token"`
	}
	raw, _ := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	_ = r.Body.Close()
	r.Body = io.NopCloser(bytes.NewReader(raw))
	_ = json.Unmarshal(raw, &body)
	if body.Email != "" {
		return body.Email
	}
	if body.Account != "" {
		return body.Account
	}
	return ""
}

func (s *Server) incDispatch() {
	if s.metrics != nil {
		s.metrics.IncDispatchRequests()
	}
}

func writeB64Proto(w http.ResponseWriter, m proto.Message) {
	raw, err := proto.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out := make([]byte, base64.StdEncoding.EncodedLen(len(raw)))
	base64.StdEncoding.Encode(out, raw)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(out)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (s *Server) dispatchURL(r *http.Request) string {
	s.mu.RLock()
	addr := s.opts.Addr
	s.mu.RUnlock()

	host := strings.TrimSpace(r.Host)
	if host == "" {
		_, port, err := parseHostPort(addr)
		if err == nil {
			host = fmt.Sprintf("127.0.0.1:%d", port)
		} else {
			host = "127.0.0.1:21000"
		}
	}
	return "http://" + host + "/query_gateway"
}

func parseHostPort(addr string) (string, int, error) {
	if addr == "" {
		return "", 0, fmt.Errorf("empty addr")
	}
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		if strings.HasPrefix(addr, ":") {
			host = "127.0.0.1"
			portStr = strings.TrimPrefix(addr, ":")
		} else {
			return "", 0, err
		}
	}
	if host == "" || host == "0.0.0.0" {
		host = "127.0.0.1"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, err
	}
	return host, port, nil
}
