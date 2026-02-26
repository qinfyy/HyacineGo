说明（对齐 LunarCore 的结构风格）：

- `internal/gameserver/server/`：服务端运行时（监听/会话/协议注册/网络收发/追踪等）。
- `internal/gameserver/server/packet/recv/`：协议处理/注册（集中存放所有 `Registry.Register(...)`）。
- `internal/gameserver/server/packet/`：cmdid、registry、编解码等。
- `internal/gameserver/game/`：游戏域模型（玩家/队伍/场景等）。
- `internal/gameserver/data/`、`internal/gameserver/database/`：资源与存储层。

开发常用脚本：

- 重新生成 cmdid：`powershell -ExecutionPolicy Bypass -File .\\scripts\\gen-cmdid.ps1`
- 重新生成 proto：`powershell -ExecutionPolicy Bypass -File .\\scripts\\gen-proto.ps1`
