# Proto layout

为避免把“协议原型（.proto）”与“编译产物（*.pb.go）”混在一起，本项目使用以下结构：

- `internal/proto/schema/StarRail.proto`：单一协议源文件（可人工编辑、可覆盖更新）
- `internal/proto/gen/`：由 `protoc-gen-go` 生成的 `*.pb.go`（不要手改）

## 重新生成

在 `tree/` 目录执行：

```powershell
.\scripts\gen-proto.ps1
```
