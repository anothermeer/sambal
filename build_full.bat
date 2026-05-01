@echo off
if not exist bin mkdir bin

echo Starting build process...

:: Linux 64-bit
echo Building for Linux (64-bit)...
set GOOS=linux
set GOARCH=amd64
go build -o bin/sambal-linux-amd64

:: Windows 64-bit
echo Building for Windows (64-bit)...
set GOOS=windows
set GOARCH=amd64
go build -o bin/sambal-windows-amd64.exe

:: macOS Intel
echo Building for macOS (Intel)...
set GOOS=darwin
set GOARCH=amd64
go build -o bin/sambal-darwin-amd64

:: macOS Apple Silicon
echo Building for macOS (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build -o bin/sambal-darwin-arm64

:: Linux ARM (Raspberry Pi)
echo Building for Linux (ARM)...
set GOOS=linux
set GOARCH=arm
go build -o bin/sambal-linux-arm

echo Done! compiled into bin/
pause
