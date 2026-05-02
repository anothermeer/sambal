#!/bin/bash

# Output directory
mkdir -p bin

echo "Starting build process..."
echo "⚠️ First time compile will be slow as fyne-cross is setting up."

# Linux 64-bit
echo "Building for Linux (64-bit)..."
fyne-cross linux -arch=amd64 -output sambal-linux-amd64 -app-id com.anothermeer.sambal

# Windows 64-bit
echo "Building for Windows (64-bit)..."
fyne-cross windows -arch=amd64 -output sambal-windows-amd64 -app-id com.anothermeer.sambal

# macOS Intel
echo "Building for macOS (Intel)..."
fyne-cross darwin -arch=amd64 -output sambal-macos-amd64 -app-id com.anothermeer.sambal

# macOS Apple Silicon
echo "Building for macOS (Apple Silicon)..."
fyne-cross darwin -arch=arm64 -output sambal-macos-arm64 -app-id com.anothermeer.sambal

# Linux ARM (Raspberry Pi)
echo "Building for Linux (ARM)..." 
fyne-cross linux -arch=arm -output sambal-linux-arm64 -app-id com.anothermeer.sambal

echo "Done! compiled into bin/"
