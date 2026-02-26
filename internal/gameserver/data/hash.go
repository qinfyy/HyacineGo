package data

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

// Hash is a compatibility wrapper for textmap/hash fields.
type Hash struct{ *big.Int }

func (h *Hash) UnmarshalJSON(b []byte) error {
	if h.Int == nil {
		h.Int = new(big.Int)
	}

	var num json.Number
	if err := json.Unmarshal(b, &num); err == nil {
		str := num.String()
		if _, ok := h.Int.SetString(str, 10); ok {
			return nil
		}
	}

	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		s = strings.TrimSpace(s)
		if _, ok := h.Int.SetString(s, 10); ok {
			return nil
		}
		return fmt.Errorf("invalid hash string: %q", s)
	}

	return fmt.Errorf("invalid hash value: %s", string(b))
}
