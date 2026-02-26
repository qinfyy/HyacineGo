param(
  [string]$Config = ".\\configs\\config.json",
  [switch]$Build
)

$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

if (-not (Test-Path $Config)) {
  throw "config not found: $Config"
}

New-Item -ItemType Directory -Force -Path ".\\configs\\players" | Out-Null
New-Item -ItemType Directory -Force -Path ".\\logs" | Out-Null
New-Item -ItemType Directory -Force -Path ".\\bin" | Out-Null

if ($Build) {
  Write-Host "building..."
  go build -o ".\\bin\\hyacine-server.exe" ".\\program\\hyacine-server"
  Write-Host "running..."
  & ".\\bin\\hyacine-server.exe" -config $Config
} else {
  Write-Host "running (go run)..."
  go run ".\\program\\hyacine-server" -config $Config
}

