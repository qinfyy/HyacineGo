package account

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

type Account struct {
	UID               string          `json:"uid"`
	Username          string          `json:"username"`
	ReservedPlayerUID int             `json:"reserved_player_uid"`
	ComboToken        string          `json:"combo_token"`
	DispatchToken     string          `json:"dispatch_token"`
	Permissions       map[string]bool `json:"permissions"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

func New(username string) (*Account, error) {
	if username == "" {
		return nil, fmt.Errorf("empty username")
	}
	now := time.Now()
	return &Account{
		UID:         newUID(),
		Username:    username,
		Permissions: map[string]bool{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (a *Account) Touch() {
	a.UpdatedAt = time.Now()
}

func (a *Account) Email() string {
	return a.Username
}

func (a *Account) GenerateComboToken() string {
	a.ComboToken = newToken()
	a.Touch()
	return a.ComboToken
}

func (a *Account) GenerateDispatchToken() string {
	a.DispatchToken = newToken()
	a.Touch()
	return a.DispatchToken
}

func newUID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func newToken() string {
	var b [32]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}
