@echo off
setlocal

set "ROOT=%~dp0.."
cd /d "%ROOT%"

set "CONFIG=.\configs\config.json"
if not "%~1"=="" set "CONFIG=%~1"

if not exist "%CONFIG%" (
  echo config not found: %CONFIG%
  exit /b 1
)

if not exist ".\configs\players" mkdir ".\configs\players" >nul 2>nul
if not exist ".\logs" mkdir ".\logs" >nul 2>nul
if not exist ".\bin" mkdir ".\bin" >nul 2>nul

echo running...
go run ".\program\hyacine-server" -config "%CONFIG%"
pause
