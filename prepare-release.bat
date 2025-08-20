@echo off
REM This script prepares the CLI for release by creating a separate go.mod file
REM that doesn't use replace directives

echo Preparing CLI for release...

REM Save the original go.mod file
copy cli\usm\go.mod cli\usm\go.mod.release.bak

REM Remove the replace directive
powershell -Command "(gc cli\usm\go.mod) -replace 'replace github.com/universal-secrets-manager/usm/core/crypto => ../../core/crypto', '' | Out-File -encoding ASCII cli\usm\go.mod"

REM Update the requirement to use the same version as the main module
cd cli\usm
go mod tidy
cd ..\..

echo CLI is ready for release. Don't forget to restore the original go.mod after building:
echo copy cli\usm\go.mod.release.bak cli\usm\go.mod