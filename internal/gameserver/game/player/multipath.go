package player

// Multi-path avatar helpers (Trailblazer + March 7th).

// TrailblazerAvatarID returns the currently selected Trailblazer avatar id for a player.
// It is a pure helper: it does not mutate the player object.
func TrailblazerAvatarID(p *Player) uint32 {
	gender := uint32(GenderWoman)
	if p != nil && p.Gender != 0 {
		gender = p.Gender
	}
	// Default Trailblazer should start at the base/physical path:
	// - Male:   TrailblazerMalePhysical
	// - Female: TrailblazerFemalePhysical
	//
	// (Hardcoding later paths like 8005/8006 tends to produce inconsistent early-game state.)
	fallback := uint32(TrailblazerFemalePhysical)
	if gender == GenderMan {
		fallback = TrailblazerMalePhysical
	}
	if p != nil && p.CurrentMultiPath != nil {
		// Multi-path Trailblazer uses a single growth key (TrailblazerGrowthAvatarID) for the current selection.
		if v := p.CurrentMultiPath[TrailblazerGrowthAvatarID]; v != 0 {
			return v
		}
		// Back-compat: older saves used TrailblazerFemalePhysical as the growth key for female Trailblazer.
		if v := p.CurrentMultiPath[TrailblazerFemalePhysical]; v != 0 {
			return v
		}
	}
	return fallback
}

// NormalizeMultiPathAvatarID maps a growth/base avatar id to the player's current selected multi-path variant.
func NormalizeMultiPathAvatarID(p *Player, id uint32) uint32 {
	if IsTrailblazerAvatarID(id) {
		return TrailblazerAvatarID(p)
	}
	if IsMarch7AvatarID(id) {
		if p != nil && p.CurrentMultiPath != nil {
			if v := p.CurrentMultiPath[MarchGrowthAvatarID]; v != 0 {
				return v
			}
		}
		return MarchPreservation
	}
	return id
}
