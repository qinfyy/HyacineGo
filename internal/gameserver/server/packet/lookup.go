package packet

import "strings"

var nameToID map[string]uint16

func init() {
	nameToID = make(map[string]uint16, len(names))
	for id, n := range names {
		nameToID[n] = id
	}
}

func IDByName(name string) (uint16, bool) {
	id, ok := nameToID[name]
	return id, ok
}

func Known(cmdID uint16) bool {
	_, ok := names[cmdID]
	return ok
}

func ScRspForCsReq(cmdID uint16) (uint16, bool) {
	n := Name(cmdID)
	if !strings.HasSuffix(n, "CsReq") {
		return 0, false
	}
	rspName := strings.TrimSuffix(n, "CsReq") + "ScRsp"
	return IDByName(rspName)
}
