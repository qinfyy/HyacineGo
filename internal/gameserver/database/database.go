package database

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"hyacine-server/internal/gameserver/game/account"
	"hyacine-server/internal/gameserver/game/player"
)

// DataBase is a simple local persistence layer for the template server.
// It stores:
// - Accounts in a single JSON file (dispatch.account_path)
// - Players as individual JSON files under a directory (gameserver.player_dir)
type DataBase struct {
	mu sync.RWMutex

	accountPath        string
	playerDir          string
	playerTemplatePath string

	accountsByUID      map[string]*account.Account
	accountsByUsername map[string]*account.Account

	playersByUID        map[uint32]*player.Player
	playersByAccountUID map[string]*player.Player

	nextPlayerUID uint32

	templateOnce sync.Once
	template     *player.Player
	templateErr  error
}

type Options struct {
	AccountPath        string
	PlayerDir          string
	PlayerTemplatePath string
}

func New(opts Options) *DataBase {
	db := &DataBase{
		accountPath:         strings.TrimSpace(opts.AccountPath),
		playerDir:           strings.TrimSpace(opts.PlayerDir),
		playerTemplatePath:  strings.TrimSpace(opts.PlayerTemplatePath),
		accountsByUID:       map[string]*account.Account{},
		accountsByUsername:  map[string]*account.Account{},
		playersByUID:        map[uint32]*player.Player{},
		playersByAccountUID: map[string]*player.Player{},
		// Start at 0 so the first auto-assigned player UID is 1.
		nextPlayerUID: 0,
	}
	if db.accountPath == "" {
		db.accountPath = "./configs/accounts.json"
	}
	if db.playerDir == "" {
		db.playerDir = "./configs/players"
	}
	if db.playerTemplatePath == "" {
		db.playerTemplatePath = "./configs/defaults/player.json"
	}
	return db
}

func (db *DataBase) AccountPath() string        { return db.accountPath }
func (db *DataBase) PlayerDir() string          { return db.playerDir }
func (db *DataBase) PlayerTemplatePath() string { return db.playerTemplatePath }

func (db *DataBase) Load() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.accountsByUID = map[string]*account.Account{}
	db.accountsByUsername = map[string]*account.Account{}
	db.playersByUID = map[uint32]*player.Player{}
	db.playersByAccountUID = map[string]*player.Player{}

	// Accounts
	raw, err := os.ReadFile(db.accountPath)
	if err == nil {
		var list []*account.Account
		if err := json.Unmarshal(raw, &list); err != nil {
			return err
		}
		for _, a := range list {
			if a == nil || a.UID == "" || a.Username == "" {
				continue
			}
			if a.Permissions == nil {
				a.Permissions = map[string]bool{}
			}
			db.accountsByUID[a.UID] = a
			db.accountsByUsername[a.Username] = a
			if a.ReservedPlayerUID > int(db.nextPlayerUID) {
				db.nextPlayerUID = uint32(a.ReservedPlayerUID)
			}
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	// Keep nextPlayerUID as a legacy hint only; allocation uses the lowest free uid starting from 1.

	return nil
}

func (db *DataBase) playerPath(uid uint32) string {
	return filepath.Join(db.playerDir, fmt.Sprintf("%d.json", uid))
}

func (db *DataBase) allocatePlayerUIDLocked() (uint32, error) {
	used := map[uint32]struct{}{}

	// Account-reserved uids.
	for _, a := range db.accountsByUID {
		if a == nil || a.ReservedPlayerUID <= 0 {
			continue
		}
		used[uint32(a.ReservedPlayerUID)] = struct{}{}
	}

	// In-memory loaded players.
	for uid := range db.playersByUID {
		used[uid] = struct{}{}
	}

	// Pick the lowest available uid starting from 1 (LC style).
	for uid := uint32(1); uid < ^uint32(0); uid++ {
		if _, ok := used[uid]; ok {
			continue
		}
		if _, err := os.Stat(db.playerPath(uid)); err == nil {
			continue
		} else if !os.IsNotExist(err) {
			return 0, err
		}
		return uid, nil
	}
	return 0, fmt.Errorf("no free uid available")
}

func (db *DataBase) loadTemplateLocked() (*player.Player, error) {
	db.templateOnce.Do(func() {
		path := strings.TrimSpace(db.playerTemplatePath)
		if path == "" {
			db.templateErr = fmt.Errorf("empty player template path")
			return
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			db.templateErr = err
			return
		}
		var p player.Player
		if err := json.Unmarshal(raw, &p); err != nil {
			db.templateErr = err
			return
		}
		db.template = &p
	})
	return db.template, db.templateErr
}

func (db *DataBase) newPlayerLocked(uid uint32, accountUID string) *player.Player {
	now := db.Now()
	if tpl, err := db.loadTemplateLocked(); err == nil && tpl != nil {
		// Deep-clone the template so runtime mutations (slices/maps) never leak back.
		raw, err := json.Marshal(tpl)
		if err == nil {
			var clone player.Player
			if err := json.Unmarshal(raw, &clone); err == nil {
				clone.UID = uid
				clone.AccountUID = accountUID
				clone.CreatedAt = now
				clone.UpdatedAt = now
				if clone.Lineups == nil {
					clone.Lineups = map[uint32]*player.LineupData{}
				}
				if clone.AvatarEnhancedIDs == nil {
					clone.AvatarEnhancedIDs = map[uint32]uint32{}
				}
				if clone.CurrentMultiPath == nil {
					clone.CurrentMultiPath = map[uint32]uint32{}
				}
				return &clone
			}
		}
	}
	p := player.New(uid, accountUID)
	p.CreatedAt = now
	p.UpdatedAt = now
	return p
}

func (db *DataBase) SaveAccounts() error {
	db.mu.RLock()
	list := make([]*account.Account, 0, len(db.accountsByUID))
	for _, a := range db.accountsByUID {
		list = append(list, a)
	}
	path := db.accountPath
	db.mu.RUnlock()

	raw, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}

func (db *DataBase) GetAccountByUID(uid string) (*account.Account, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	a, ok := db.accountsByUID[uid]
	return a, ok
}

func (db *DataBase) GetAccountByUsername(username string) (*account.Account, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	a, ok := db.accountsByUsername[username]
	return a, ok
}

func (db *DataBase) GetOrCreateAccount(username string) (*account.Account, bool, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, false, fmt.Errorf("empty username")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if a, ok := db.accountsByUsername[username]; ok {
		return a, false, nil
	}
	a, err := account.New(username)
	if err != nil {
		return nil, false, err
	}
	a.Touch()
	a.GenerateDispatchToken()
	a.GenerateComboToken()
	db.accountsByUID[a.UID] = a
	db.accountsByUsername[a.Username] = a
	return a, true, nil
}

func (db *DataBase) EnsurePlayerForAccount(accountUID string, reservedPlayerUID uint32) (*player.Player, bool, error) {
	accountUID = strings.TrimSpace(accountUID)
	if accountUID == "" {
		return nil, false, fmt.Errorf("empty account uid")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if p, ok := db.playersByAccountUID[accountUID]; ok {
		return p, false, nil
	}

	uid := reservedPlayerUID
	if uid == 0 {
		allocated, err := db.allocatePlayerUIDLocked()
		if err != nil {
			return nil, false, err
		}
		uid = allocated
	}

	// Try load existing file.
	if loaded, err := db.loadPlayerLocked(uid); err == nil && loaded != nil {
		// Older player files may miss uid/account_uid; normalize to the requested identity.
		if loaded.UID == 0 {
			loaded.UID = uid
		}
		if strings.TrimSpace(loaded.AccountUID) == "" {
			loaded.AccountUID = accountUID
		}
		db.playersByUID[loaded.UID] = loaded
		db.playersByAccountUID[loaded.AccountUID] = loaded
		return loaded, false, nil
	} else if err != nil && !os.IsNotExist(err) {
		// Don't silently overwrite an existing (but invalid) player file with template defaults.
		return nil, false, err
	}

	p := db.newPlayerLocked(uid, accountUID)
	db.playersByUID[p.UID] = p
	db.playersByAccountUID[p.AccountUID] = p
	if err := db.savePlayerLocked(p); err != nil {
		return nil, false, err
	}
	return p, true, nil
}

func (db *DataBase) SavePlayer(p *player.Player) error {
	if p == nil {
		return nil
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.savePlayerLocked(p)
}

func (db *DataBase) GetPlayerByUID(uid uint32) (*player.Player, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	if p, ok := db.playersByUID[uid]; ok && p != nil {
		return p, true
	}
	loaded, err := db.loadPlayerLocked(uid)
	if err != nil || loaded == nil {
		return nil, false
	}
	db.playersByUID[loaded.UID] = loaded
	db.playersByAccountUID[loaded.AccountUID] = loaded
	return loaded, true
}

func (db *DataBase) GetPlayerByAccountUID(accountUID string) (*player.Player, bool) {
	accountUID = strings.TrimSpace(accountUID)
	if accountUID == "" {
		return nil, false
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	if p, ok := db.playersByAccountUID[accountUID]; ok && p != nil {
		return p, true
	}
	return nil, false
}

func (db *DataBase) savePlayerLocked(p *player.Player) error {
	if err := os.MkdirAll(db.playerDir, 0o755); err != nil {
		return err
	}
	p.Touch()
	raw, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(db.playerPath(p.UID), raw, 0o644)
}

func (db *DataBase) loadPlayerLocked(uid uint32) (*player.Player, error) {
	raw, err := os.ReadFile(db.playerPath(uid))
	if err != nil {
		return nil, err
	}
	var p player.Player
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, err
	}
	if p.UID == 0 {
		p.UID = uid
	}
	return &p, nil
}

func newToken() string {
	var b [32]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func (db *DataBase) Now() time.Time { return time.Now() }
