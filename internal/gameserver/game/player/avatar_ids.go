package player

// Named avatar ids for multi-path characters.
//
// Keep these as plain uint32 constants so call sites don't need to remember magic numbers.

const (
	// Trailblazer (multi-path, gendered).
	TrailblazerGrowthAvatarID uint32 = 8001

	TrailblazerMalePhysical       uint32 = 8001
	TrailblazerFemalePhysical     uint32 = 8002
	TrailblazerMalePreservation   uint32 = 8003
	TrailblazerFemalePreservation uint32 = 8004
	TrailblazerMaleHarmony        uint32 = 8005
	TrailblazerFemaleHarmony      uint32 = 8006
	TrailblazerMaleRemembrance    uint32 = 8007
	TrailblazerFemaleRemembrance  uint32 = 8008

	// March 7th (multi-path, non-gendered).
	MarchGrowthAvatarID uint32 = 1001

	MarchPreservation uint32 = 1001
	MarchHunt         uint32 = 1224
)

func IsTrailblazerMaleAvatarID(id uint32) bool {
	switch id {
	case TrailblazerMalePhysical, TrailblazerMalePreservation, TrailblazerMaleHarmony, TrailblazerMaleRemembrance:
		return true
	default:
		return false
	}
}

func IsTrailblazerFemaleAvatarID(id uint32) bool {
	switch id {
	case TrailblazerFemalePhysical, TrailblazerFemalePreservation, TrailblazerFemaleHarmony, TrailblazerFemaleRemembrance:
		return true
	default:
		return false
	}
}

func IsTrailblazerAvatarID(id uint32) bool {
	return IsTrailblazerMaleAvatarID(id) || IsTrailblazerFemaleAvatarID(id)
}

func IsMarch7AvatarID(id uint32) bool {
	switch id {
	case MarchPreservation, MarchHunt:
		return true
	default:
		return false
	}
}
