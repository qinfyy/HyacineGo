package server

import (
	"sync"
)

var (
	mainMissionIDsOnce sync.Once
	mainMissionIDs     []uint32

	subMissionIDsOnce sync.Once
	subMissionIDs     []uint32

	tutorialIDsOnce      sync.Once
	tutorialIDs          []uint32
	tutorialGuideIDsOnce sync.Once
	tutorialGuideIDs     []uint32
)

func (s *Server) loadAllMainMissionIDs() []uint32 {
	mainMissionIDsOnce.Do(func() {
		if s == nil || s.data == nil || s.data.Loader() == nil {
			mainMissionIDs = nil
			return
		}
		var raw []struct {
			MainMissionID uint32 `json:"MainMissionID"`
		}
		if err := s.data.Loader().LoadExcelJSON("MainMission", &raw); err != nil {
			mainMissionIDs = nil
			return
		}
		ids := make([]uint32, 0, len(raw))
		seen := make(map[uint32]struct{}, len(raw))
		for _, r := range raw {
			if r.MainMissionID == 0 {
				continue
			}
			if _, ok := seen[r.MainMissionID]; ok {
				continue
			}
			seen[r.MainMissionID] = struct{}{}
			ids = append(ids, r.MainMissionID)
		}
		mainMissionIDs = ids
	})
	return mainMissionIDs
}

func (s *Server) loadAllSubMissionIDs() []uint32 {
	subMissionIDsOnce.Do(func() {
		if s == nil || s.data == nil || s.data.Loader() == nil {
			subMissionIDs = nil
			return
		}
		var raw []struct {
			SubMissionID uint32 `json:"SubMissionID"`
		}
		if err := s.data.Loader().LoadExcelJSON("SubMission", &raw); err != nil {
			subMissionIDs = nil
			return
		}
		ids := make([]uint32, 0, len(raw))
		seen := make(map[uint32]struct{}, len(raw))
		for _, r := range raw {
			if r.SubMissionID == 0 {
				continue
			}
			if _, ok := seen[r.SubMissionID]; ok {
				continue
			}
			seen[r.SubMissionID] = struct{}{}
			ids = append(ids, r.SubMissionID)
		}
		subMissionIDs = ids
	})
	return subMissionIDs
}

func (s *Server) loadAllTutorialIDs() []uint32 {
	tutorialIDsOnce.Do(func() {
		if s == nil || s.data == nil || s.data.Loader() == nil {
			tutorialIDs = nil
			return
		}
		var raw []struct {
			TutorialID uint32 `json:"TutorialID"`
		}
		if err := s.data.Loader().LoadExcelJSON("TutorialData", &raw); err != nil {
			tutorialIDs = nil
			return
		}
		ids := make([]uint32, 0, len(raw))
		seen := make(map[uint32]struct{}, len(raw))
		for _, r := range raw {
			if r.TutorialID == 0 {
				continue
			}
			if _, ok := seen[r.TutorialID]; ok {
				continue
			}
			seen[r.TutorialID] = struct{}{}
			ids = append(ids, r.TutorialID)
		}
		tutorialIDs = ids
	})
	return tutorialIDs
}

func (s *Server) loadAllTutorialGuideIDs() []uint32 {
	tutorialGuideIDsOnce.Do(func() {
		if s == nil || s.data == nil || s.data.Loader() == nil {
			tutorialGuideIDs = nil
			return
		}
		var raw []struct {
			TutorialGuideID uint32 `json:"TutorialGuideID"`
		}
		if err := s.data.Loader().LoadExcelJSON("TutorialGuideData", &raw); err != nil {
			tutorialGuideIDs = nil
			return
		}
		ids := make([]uint32, 0, len(raw))
		seen := make(map[uint32]struct{}, len(raw))
		for _, r := range raw {
			if r.TutorialGuideID == 0 {
				continue
			}
			if _, ok := seen[r.TutorialGuideID]; ok {
				continue
			}
			seen[r.TutorialGuideID] = struct{}{}
			ids = append(ids, r.TutorialGuideID)
		}
		tutorialGuideIDs = ids
	})
	return tutorialGuideIDs
}
