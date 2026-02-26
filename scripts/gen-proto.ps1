param(
  [string]$ProtoDir = ".\\internal\\proto\\schema",
  [string]$OutDir = ".\\internal\\proto\\gen"
)

$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

if (-not (Test-Path $ProtoDir)) {
    throw "proto directory not found: $ProtoDir"
}

New-Item -ItemType Directory -Force -Path $OutDir | Out-Null

$protoFiles = Get-ChildItem -Path $ProtoDir -Filter *.proto -File
if ($protoFiles.Count -eq 0) {
    throw "no .proto files found in $ProtoDir"
}

Write-Host "generating protobuf Go files from $ProtoDir..."

foreach ($proto in $protoFiles) {
    $protoAbs = $proto.FullName
    $protoName = $proto.Name
    $outImport = "hyacine-server/" + (($OutDir -replace "^\.\\" , "") -replace "\\", "/")

    Push-Location $proto.DirectoryName
    protoc -I . --go_out=$root --go_opt=module=hyacine-server --go_opt="M$protoName=$outImport" $protoName
    if ($LASTEXITCODE -ne 0) { throw "protoc failed for $protoName" }
    Pop-Location
}

Write-Host "done: $OutDir"
