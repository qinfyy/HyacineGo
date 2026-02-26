package observability

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

type Metrics struct {
	dispatchRequests atomic.Uint64
	gameConnections  atomic.Uint64
	gameFrames       atomic.Uint64
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) IncDispatchRequests() { m.dispatchRequests.Add(1) }
func (m *Metrics) IncGameConnections()  { m.gameConnections.Add(1) }
func (m *Metrics) IncGameFrames()       { m.gameFrames.Add(1) }

func (m *Metrics) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"dispatch_requests_total": m.dispatchRequests.Load(),
			"game_connections_total":  m.gameConnections.Load(),
			"game_frames_total":       m.gameFrames.Load(),
		})
	})
}
