package gear

import "hyacine-server/internal/gameserver/freesrdata"

const (
	FreesrEquipmentUIDBase uint32 = 2_000_000_000
	FreesrRelicUIDBase     uint32 = 3_000_000_000
)

func FreesrEquipmentUniqueID(lc freesrdata.LightconePreset) uint32 {
	if lc.InternalUID != 0 {
		return lc.InternalUID
	}
	if lc.EquipAvatar == 0 || lc.ItemID == 0 {
		return 0
	}
	// Keep stable and within uint32. Most lightcone item_ids fit in <100000.
	return FreesrEquipmentUIDBase + lc.EquipAvatar*100_000 + (lc.ItemID % 100_000)
}

func FreesrRelicUniqueID(r freesrdata.RelicPreset) uint32 {
	if r.InternalUID != 0 {
		return r.InternalUID
	}
	if r.EquipAvatar == 0 || r.RelicID == 0 {
		return 0
	}
	return FreesrRelicUIDBase + r.EquipAvatar*100_000 + (r.RelicID % 100_000)
}

func RelicSlotTypeFromTID(tid uint32) uint32 {
	// Empirically FreeSR relic ids end with 1..6 for the slot.
	t := tid % 10
	if t >= 1 && t <= 6 {
		return t
	}
	return 0
}
