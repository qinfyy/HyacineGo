package data

import (
	"encoding/json"
	"fmt"
	"math"
)

// HashU64 is a compatibility wrapper for textmap/hash fields.
//
// Some resource dumps encode 32-bit hashes as signed integers (e.g. -712088400).
// Official servers treat these as unsigned (uint32) hashes. This type accepts
// both signed/unsigned JSON numbers and normalizes them into an unsigned value.
type HashU64 uint64

func (h *HashU64) UnmarshalJSON(b []byte) error {
	// Try number first.
	var n int64
	if err := json.Unmarshal(b, &n); err == nil {
		if n < 0 {
			if n < math.MinInt32 {
				return fmt.Errorf("hash out of int32 range: %d", n)
			}
			*h = HashU64(uint64(uint32(int32(n))))
			return nil
		}
		*h = HashU64(uint64(n))
		return nil
	}

	// Some dumps may encode hashes as strings.
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		var u uint64
		if _, err := fmt.Sscanf(s, "%d", &u); err == nil {
			*h = HashU64(u)
			return nil
		}
		return fmt.Errorf("invalid hash string: %q", s)
	}

	return fmt.Errorf("invalid hash: %s", string(b))
}
