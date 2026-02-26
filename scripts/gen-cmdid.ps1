param(
  [string]$In = ".\\scripts\\cmdids.txt",
  [string]$Out = ".\\internal\\gameserver\\server\\packet\\cmdid.go"
)

$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

if (-not (Test-Path $In)) {
  throw "cmdid source not found: $In"
}

$rx = [regex]'^\s*Cmd(?<name>[A-Za-z0-9_]+)\s*=\s*(?<id>\d+)\s*,?\s*$'

$byId = @{}
$byName = @{}

foreach ($line in Get-Content $In) {
  $m = $rx.Match($line)
  if (-not $m.Success) { continue }
  $name = $m.Groups['name'].Value
  $id = [int]$m.Groups['id'].Value

  if ($id -lt 0 -or $id -gt 65535) {
    throw "cmd id out of uint16 range: $id ($name)"
  }

  if ($byName.ContainsKey($name)) { continue }
  if ($byId.ContainsKey($id)) { continue }

  $byName[$name] = $id
  $byId[$id] = $name
}

if ($byId.Count -eq 0) {
  throw "no cmd ids parsed from: $In"
}

$ordered = $byId.GetEnumerator() | Sort-Object Name | ForEach-Object {
  [pscustomobject]@{ Id = [int]$_.Key; Name = [string]$_.Value }
}

$nl = "`r`n"
$sb = New-Object System.Text.StringBuilder

[void]$sb.Append("package packet$nl$nl")
[void]$sb.Append("// Code generated from scripts/cmdids.txt; DO NOT EDIT.$nl$nl")
[void]$sb.Append("const ($nl")
[void]$sb.Append("`tNONE uint16 = 0$nl")
foreach ($it in $ordered) {
  [void]$sb.Append("`t$($it.Name) uint16 = $($it.Id)$nl")
}
[void]$sb.Append(")$nl$nl")

[void]$sb.Append("var names = map[uint16]string{$nl")
[void]$sb.Append("`tNONE: `"NONE`",$nl")
foreach ($it in $ordered) {
  [void]$sb.Append("`t$($it.Name): `"$($it.Name)`",$nl")
}
[void]$sb.Append("}$nl$nl")

[void]$sb.Append("func Name(cmdID uint16) string {$nl")
[void]$sb.Append("`tif s, ok := names[cmdID]; ok {$nl")
[void]$sb.Append("`t`treturn s$nl")
[void]$sb.Append("`t}$nl")
[void]$sb.Append("`treturn `"Cmd(`" + itoa(uint32(cmdID)) + `")`"$nl")
[void]$sb.Append("}$nl$nl")

[void]$sb.Append("func itoa(v uint32) string {$nl")
[void]$sb.Append("`tif v == 0 {$nl")
[void]$sb.Append("`t`treturn `"0`"$nl")
[void]$sb.Append("`t}$nl")
[void]$sb.Append("`tvar buf [10]byte$nl")
[void]$sb.Append("`ti := len(buf)$nl")
[void]$sb.Append("`tfor v > 0 {$nl")
[void]$sb.Append("`t`ti--$nl")
[void]$sb.Append("`t`tbuf[i] = byte('0' + v%10)$nl")
[void]$sb.Append("`t`tv /= 10$nl")
[void]$sb.Append("`t}$nl")
[void]$sb.Append("`treturn string(buf[i:])$nl")
[void]$sb.Append("}$nl")

$outDir = Split-Path -Parent $Out
if ($outDir) {
  New-Item -ItemType Directory -Force -Path $outDir | Out-Null
}

Set-Content -Path $Out -Value $sb.ToString() -NoNewline

$gofmt = Get-Command gofmt -ErrorAction SilentlyContinue
if ($null -ne $gofmt) {
  & $gofmt.Path -w $Out
}

Write-Host "generated: $Out"
