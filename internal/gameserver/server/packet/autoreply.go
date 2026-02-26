package packet

import (
	"log/slog"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var (
	autoReplyOnce sync.Once
	autoReplyMu   sync.RWMutex
	autoReplyByGo map[string]protoreflect.MessageType // short Go/Proto name (e.g. "PlayerLoginScRsp") -> type
)

func initAutoReplyTypes() {
	autoReplyByGo = map[string]protoreflect.MessageType{}
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		name := string(mt.Descriptor().Name())
		autoReplyByGo[name] = mt
		return true
	})
}

func newMessageByShortName(short string) (proto.Message, bool) {
	autoReplyOnce.Do(initAutoReplyTypes)
	autoReplyMu.RLock()
	mt, ok := autoReplyByGo[short]
	autoReplyMu.RUnlock()
	if !ok {
		return nil, false
	}
	return mt.New().Interface(), true
}

// AutoReplyScRsp builds a best-effort ScRsp protobuf message for a CsReq cmd id.
// It sets Retcode=0 if such field exists, otherwise leaves defaults.
func AutoReplyScRsp(csReqCmdID uint16) (rspCmdID uint16, msg proto.Message, ok bool) {
	rspCmdID, ok = ScRspForCsReq(csReqCmdID)
	if !ok {
		return 0, nil, false
	}

	rspName := Name(rspCmdID)
	if rspName == "" || rspName == "Unknown" {
		return 0, nil, false
	}

	m, ok := newMessageByShortName(rspName)
	if !ok {
		// Some projects generate odd/hashed filenames; still the message type should exist if compiled.
		// If it doesn't, keep the legacy empty-packet fallback.
		slog.Debug("autoreply type not found", "name", rspName)
		return 0, nil, false
	}

	setRetcodeZero(m)
	return rspCmdID, m, true
}

func setRetcodeZero(m proto.Message) {
	pm := m.ProtoReflect()
	fd := pm.Descriptor().Fields().ByName("retcode")
	if fd == nil {
		// Some messages use "Retcode" in source, but protoreflect uses proto field name.
		// Keep a tiny fallback for unusual generators.
		fd = pm.Descriptor().Fields().ByName(protoreflect.Name(strings.ToLower("Retcode")))
	}
	if fd == nil {
		return
	}
	switch fd.Kind() {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		pm.Set(fd, protoreflect.ValueOfInt32(0))
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		pm.Set(fd, protoreflect.ValueOfInt64(0))
	}
}
