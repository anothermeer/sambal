#!/usr/bin/env bash

set -e

MACOSSDK=""

# Parse arguments
while [[ $# -gt 0 ]]; do
  case "$1" in
    --macossdk)
      MACOSSDK="$2"
      shift 2
      ;;
    *)
      echo "Unknown argument: $1"
      exit 1
      ;;
  esac
done

# Create bin directory if not exists
mkdir -p bin

echo "Starting build process..."
echo "[!] First time compile will be slow as fyne-cross is setting up."

export GOTOOLCHAIN=auto
export FYNE_CROSS_GO_VERSION=1.26.2

# Optional macOS SDK
if [[ -n "$MACOSSDK" ]]; then
  echo "Using macOS SDK: $MACOSSDK"
  export MACOS_SDK_PATH="$MACOSSDK"
fi

# Check Docker
if ! docker info >/dev/null 2>&1; then
  echo "Docker is not running. Please start Docker first."
  read -p "Press enter to exit"
  exit 1
fi

echo "Building for Linux (64-bit)..."
fyne-cross linux -arch=amd64 -output sambal-linux-amd64 -app-id com.anothermeer.sambal

echo "Building for Windows (64-bit)..."
fyne-cross windows -arch=amd64 -output sambal-windows-amd64 -app-id com.anothermeer.sambal

echo "Building for macOS (Intel)..."
fyne-cross darwin -arch=amd64 -output sambal-macos-amd64 -app-id com.anothermeer.sambal

echo "Building for macOS (Apple Silicon)..."
fyne-cross darwin -arch=arm64 -output sambal-macos-arm64 -app-id com.anothermeer.sambal

echo "Building for Linux (ARM)..."
fyne-cross linux -arch=arm -output sambal-linux-arm64 -app-id com.anothermeer.sambal

echo "Copying outputs to bin/..."
cp -r fyne-cross/dist/* bin/

echo "All done! You should be seeing 5 binaries in the bin folder"
read -p "Press enter to exit"