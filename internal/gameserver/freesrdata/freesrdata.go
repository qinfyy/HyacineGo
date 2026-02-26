package freesrdata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Data struct {
	Avatars      map[uint32]AvatarPreset
	Lightcones   []LightconePreset
	Relics       []RelicPreset
	BattleConfig *BattleConfig
	Key          string
}

type AvatarPreset struct {
	AvatarID   uint32   `json:"avatar_id"`
	Level      uint32   `json:"level"`
	Promotion  uint32   `json:"promotion"`
	SpValue    uint32   `json:"sp_value"`
	SpMax      uint32   `json:"sp_max"`
	EnhancedID uint32   `json:"enhanced_id"`
	Techniques []uint32 `json:"techniques"`
	Data       struct {
		Rank   uint32         `json:"rank"`
		Skills map[string]int `json:"skills"`
	} `json:"data"`
}

type LightconePreset struct {
	InternalUID uint32 `json:"internal_uid"`
	EquipAvatar uint32 `json:"equip_avatar"`
	ItemID      uint32 `json:"item_id"`
	Promotion   uint32 `json:"promotion"`
	Rank        uint32 `json:"rank"`
	Level       uint32 `json:"level"`
}

type RelicSubAffix struct {
	Count      uint32 `json:"count"`
	Step       uint32 `json:"step"`
	SubAffixID uint32 `json:"sub_affix_id"`
}

type RelicPreset struct {
	InternalUID  uint32          `json:"internal_uid"`
	EquipAvatar  uint32          `json:"equip_avatar"`
	Level        uint32          `json:"level"`
	MainAffixID  uint32          `json:"main_affix_id"`
	RelicID      uint32          `json:"relic_id"`
	RelicSetID   uint32          `json:"relic_set_id"`
	SubAffixes   []RelicSubAffix `json:"sub_affixes"`
	AnchorPosX   float32         `json:"pos_x,omitempty"`
	AnchorPosY   float32         `json:"pos_y,omitempty"`
	AnchorPosZ   float32         `json:"pos_z,omitempty"`
	AnchorPosInt uint32          `json:"pos_int,omitempty"`
}

type BattleMonster struct {
	MonsterID uint32 `json:"monster_id"`
	Amount    uint32 `json:"amount"`
	Level     uint32 `json:"level"`
}

type BattleBlessing struct {
	Level uint32 `json:"level"`
	ID    uint32 `json:"id"`
}

type BattleConfig struct {
	BattleType      uint32            `json:"battle_type"`
	CycleCount      uint32            `json:"cycle_count"`
	StageID         uint32            `json:"stage_id"`
	PathResonanceID uint32            `json:"path_resonance_id"`
	Monsters        [][]BattleMonster `json:"monsters"`
	Blessings       []BattleBlessing  `json:"blessings"`
	CustomStats     []json.RawMessage `json:"custom_stats"`
}

type rootFull struct {
	Avatars      json.RawMessage   `json:"avatars"`
	Lightcones   []LightconePreset `json:"lightcones"`
	Relics       []RelicPreset     `json:"relics"`
	BattleConfig *BattleConfig     `json:"battle_config"`
	Key          string            `json:"key"`
}

func newData() *Data {
	return &Data{
		Avatars:    map[uint32]AvatarPreset{},
		Lightcones: nil,
		Relics:     nil,
	}
}

func LoadDir(dir string) (*Data, error) {
	return Load(dir)
}

// Load reads FreeSR-like presets from a file or directory.
// Directory mode merges all JSON files under the directory (best-effort).
func Load(path string) (*Data, error) {
	if strings.TrimSpace(path) == "" {
		return newData(), nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	out := newData()
	if !info.IsDir() {
		if err := loadFileInto(out, path); err != nil {
			return nil, err
		}
		return out, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(ent.Name())) != ".json" {
			continue
		}
		fp := filepath.Join(path, ent.Name())
		if err := loadFileInto(out, fp); err != nil {
			// ignore individual files; caller will decide strictness
			continue
		}
	}
	return out, nil
}

func loadFileInto(out *Data, path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// FreeSR format: {"avatars": {...}, "lightcones": [...], "relics": [...], ...}
	var full rootFull
	if err := json.Unmarshal(raw, &full); err == nil {
		if full.Avatars != nil || len(full.Lightcones) > 0 || len(full.Relics) > 0 || full.BattleConfig != nil || strings.TrimSpace(full.Key) != "" {
			if full.Avatars != nil {
				if err := mergeAvatars(out, full.Avatars); err != nil {
					return err
				}
			}
			if len(full.Lightcones) > 0 {
				out.Lightcones = append(out.Lightcones, full.Lightcones...)
			}
			if len(full.Relics) > 0 {
				out.Relics = append(out.Relics, full.Relics...)
			}
			if full.BattleConfig != nil {
				out.BattleConfig = full.BattleConfig
			}
			if strings.TrimSpace(full.Key) != "" {
				out.Key = full.Key
			}
			return nil
		}
	}

	// Fallback format: [{"avatar_id":1001,...}, ...]
	var list []AvatarPreset
	if err := json.Unmarshal(raw, &list); err != nil {
		// Legacy map-only format: {"avatars": {"1001": {...}, ...}}
		var r struct {
			Avatars map[string]AvatarPreset `json:"avatars"`
		}
		if err2 := json.Unmarshal(raw, &r); err2 == nil && len(r.Avatars) > 0 {
			for k, v := range r.Avatars {
				id := v.AvatarID
				if id == 0 {
					parsed, _ := strconv.ParseUint(k, 10, 32)
					id = uint32(parsed)
				}
				if id == 0 {
					continue
				}
				v.AvatarID = id
				out.Avatars[id] = v
			}
			return nil
		}
		return err
	}
	for _, v := range list {
		if v.AvatarID == 0 {
			continue
		}
		out.Avatars[v.AvatarID] = v
	}
	return nil
}

func mergeAvatars(out *Data, raw json.RawMessage) error {
	if out == nil {
		return nil
	}
	raw = json.RawMessage(strings.TrimSpace(string(raw)))
	if len(raw) == 0 {
		return nil
	}

	// Common format: {"1001": {...}, ...}
	var m map[string]AvatarPreset
	if err := json.Unmarshal(raw, &m); err == nil && len(m) > 0 {
		for k, v := range m {
			id := v.AvatarID
			if id == 0 {
				parsed, _ := strconv.ParseUint(k, 10, 32)
				id = uint32(parsed)
			}
			if id == 0 {
				continue
			}
			v.AvatarID = id
			out.Avatars[id] = v
		}
		return nil
	}

	// Alternative format: [{"avatar_id":1001,...}, ...]
	var list []AvatarPreset
	if err := json.Unmarshal(raw, &list); err != nil {
		return err
	}
	for _, v := range list {
		if v.AvatarID == 0 {
			continue
		}
		out.Avatars[v.AvatarID] = v
	}
	return nil
}

func (d *Data) AvatarCount() int {
	if d == nil {
		return 0
	}
	return len(d.Avatars)
}

func (d *Data) LightconeCount() int {
	if d == nil {
		return 0
	}
	return len(d.Lightcones)
}

func (d *Data) RelicCount() int {
	if d == nil {
		return 0
	}
	return len(d.Relics)
}

func (d *Data) AvatarIDs() []uint32 {
	if d == nil || len(d.Avatars) == 0 {
		return nil
	}
	out := make([]uint32, 0, len(d.Avatars))
	for id := range d.Avatars {
		out = append(out, id)
	}
	// deterministic order without importing sort (tiny)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j] < out[i] {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

func (d *Data) GetAvatar(id uint32) (AvatarPreset, bool) {
	if d == nil {
		return AvatarPreset{}, false
	}
	v, ok := d.Avatars[id]
	return v, ok
}

func (d *Data) MustHaveAvatars() error {
	if d == nil || len(d.Avatars) == 0 {
		return fmt.Errorf("no freesr avatars loaded")
	}
	return nil
}

func (d *Data) LightconesForAvatar(avatarID uint32) []LightconePreset {
	if d == nil || avatarID == 0 || len(d.Lightcones) == 0 {
		return nil
	}
	out := make([]LightconePreset, 0, 1)
	for _, lc := range d.Lightcones {
		if lc.EquipAvatar == avatarID {
			out = append(out, lc)
		}
	}
	return out
}

func (d *Data) RelicsForAvatar(avatarID uint32) []RelicPreset {
	if d == nil || avatarID == 0 || len(d.Relics) == 0 {
		return nil
	}
	out := make([]RelicPreset, 0, 6)
	for _, r := range d.Relics {
		if r.EquipAvatar == avatarID {
			out = append(out, r)
		}
	}
	return out
}
