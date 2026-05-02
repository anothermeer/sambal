@echo off
if not exist bin mkdir bin

echo Starting build process...
echo [!] First time compile will be slow as fyne-cross is setting up.

set GOTOOLCHAIN=auto

docker info >nul 2>&1
if errorlevel 1 (
    echo Docker is not running. Please start Docker Desktop first.
    pause
    exit /b
)

:: Linux 64-bit
echo Building for Linux (64-bit)...
fyne-cross linux -arch=amd64 -output sambal-linux-amd64 -app-id com.anothermeer.sambal

:: Windows 64-bit
echo Building for Windows (64-bit)...
fyne-cross windows -arch=amd64 -output sambal-windows-amd64 -app-id com.anothermeer.sambal

:: macOS Intel
echo Building for macOS (Intel)...
fyne-cross darwin -arch=amd64 -output sambal-macos-amd64 -app-id com.anothermeer.sambal

:: macOS Apple Silicon
echo Building for macOS (Apple Silicon)...
fyne-cross darwin -arch=arm64 -output sambal-macos-arm64 -app-id com.anothermeer.sambal

:: Linux ARM (Raspberry Pi)
echo Building for Linux (ARM)...
fyne-cross linux -arch=arm -output sambal-linux-arm64 -app-id com.anothermeer.sambal

echo Copying outputs to bin/...
xcopy /E /I /Y fyne-cross\dist\* bin\

echo All done! You should be seeing 5 binaries in the bin folder
pause
