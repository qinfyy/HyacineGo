param(
  [string]$Proto = ".\\internal\\proto\\schema\\StarRail.proto",
  [string]$OutDir = ".\\internal\\proto\\gen"
)

$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

if (-not (Test-Path $Proto)) {
  throw "proto not found: $Proto"
}

New-Item -ItemType Directory -Force -Path $OutDir | Out-Null

$protoAbs = (Resolve-Path $Proto).Path
$protoDir = Split-Path -Parent $protoAbs
$protoName = Split-Path -Leaf $protoAbs
$outImport = "hyacine-server/" + (($OutDir -replace "^\\.\\\\", "") -replace "\\\\", "/")

Write-Host "generating protobuf Go files..."

# Use module=hyacine-server to avoid creating a nested "hyacine-server/" output folder.
Push-Location $protoDir
protoc -I . --go_out=$root --go_opt=module=hyacine-server --go_opt="M$protoName=$outImport" $protoName
if ($LASTEXITCODE -ne 0) { throw "protoc failed" }
Pop-Location

Write-Host "done: $OutDir"
