param (
    [string]$macossdk
)

# Parse --macossdk manually (supports: --macossdk path)
for ($i = 0; $i -lt $args.Length; $i++) {
    if ($args[$i] -eq "--macossdk" -and $i + 1 -lt $args.Length) {
        $macossdk = $args[$i + 1]
    }
}

# Create bin directory if not exists
if (-not (Test-Path "bin")) {
    New-Item -ItemType Directory -Path "bin" | Out-Null
}

Write-Host "Starting build process..."
Write-Host "[!] First time compile will be slow as fyne-cross is setting up."

$env:GOTOOLCHAIN = "auto"
$env:FYNE_CROSS_GO_VERSION = "1.26.2"

# Optional: use macOS SDK if provided
if ($macossdk) {
    Write-Host "Using macOS SDK: $macossdk"
    $env:MACOS_SDK_PATH = $macossdk
}

# Check Docker
docker info *> $null
if ($LASTEXITCODE -ne 0) {
    Write-Host "Docker is not running. Please start Docker Desktop first."
    Read-Host "Press Enter to exit"
    exit 1
}

Write-Host "Building for Linux (64-bit)..."
fyne-cross linux -arch=amd64 -output sambal-linux-amd64 -app-id com.anothermeer.sambal

Write-Host "Building for Windows (64-bit)..."
fyne-cross windows -arch=amd64 -output sambal-windows-amd64 -app-id com.anothermeer.sambal

# I'm not going to build MacOS builds for now... such headache
#Write-Host "Building for macOS (Intel)..."
#fyne-cross darwin -arch=amd64 -output sambal-macos-amd64 -app-id com.anothermeer.sambal

#Write-Host "Building for macOS (Apple Silicon)..."
#fyne-cross darwin -arch=arm64 -output sambal-macos-arm64 -app-id com.anothermeer.sambal

Write-Host "Building for Linux (ARM)..."
fyne-cross linux -arch=arm -output sambal-linux-arm64 -app-id com.anothermeer.sambal

Write-Host "Copying outputs to bin/..."
Copy-Item -Recurse -Force "fyne-cross/dist/*" "bin/"

Write-Host "All done! You should be seeing 3 binaries in the bin folder"
Read-Host "Press Enter to exit"