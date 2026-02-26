package packet

import (
	"log/slog"
	"time"

	"google.golang.org/protobuf/proto"
)

type SessionState int

const (
	StateWaitingForToken SessionState = iota
	StateWaitingForLogin
	StateActive
)

type Context struct {
	Session   *Session
	Send      func(cmdID uint16, msg proto.Message)
	SendEmpty func(cmdID uint16)
	Now       func() time.Time
	Log       *slog.Logger
}

type HandlerFunc func(ctx Context, cmdID uint16, payload []byte)

type Handler struct {
	Name      string
	MinState  SessionState
	AllowWhen func(state SessionState) bool
	Fn        HandlerFunc
}

type Registry struct {
	handlers map[uint16]Handler
}

func NewRegistry() *Registry {
	return &Registry{handlers: map[uint16]Handler{}}
}

func (r *Registry) Register(cmdID uint16, h Handler) {
	if h.Name == "" {
		h.Name = Name(cmdID)
	}
	r.handlers[cmdID] = h
}

func (r *Registry) Handle(ctx Context, cmdID uint16, payload []byte) bool {
	h, ok := r.handlers[cmdID]
	if !ok || h.Fn == nil {
		return false
	}

	state := ctx.Session.State
	if h.AllowWhen != nil {
		if !h.AllowWhen(state) {
			return true
		}
	} else if state < h.MinState {
		return true
	}

	h.Fn(ctx, cmdID, payload)
	return true
}

type Session struct {
	State SessionState

	AccountUID string
	UID        uint32

	LastActive time.Time

	// HasEnteredWorld marks whether we already pushed an initial scene entry for this session.
	HasEnteredWorld bool

	// TrackMainMissionID mirrors the client's tracked main mission id (for UI pointers / map gating).
	TrackMainMissionID uint32
}

func (s SessionState) String() string {
	switch s {
	case StateWaitingForToken:
		return "WAITING_FOR_TOKEN"
	case StateWaitingForLogin:
		return "WAITING_FOR_LOGIN"
	case StateActive:
		return "ACTIVE"
	default:
		return "UNKNOWN"
	}
}
