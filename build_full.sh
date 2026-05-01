#!/bin/bash

# Output directory
mkdir -p bin

echo "Starting build process..."

# Linux 64-bit
echo "Building for Linux (64-bit)..."
GOOS=linux GOARCH=amd64 go build -o bin/sambal-linux-amd64

# Windows 64-bit
echo "Building for Windows (64-bit)..."
GOOS=windows GOARCH=amd64 go build -o bin/sambal-windows-amd64.exe

# macOS Intel
echo "Building for macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o bin/sambal-darwin-amd64

# macOS Apple Silicon
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o bin/sambal-darwin-arm64

# Linux ARM (Raspberry Pi)
echo "Building for Linux (ARM)..."
GOOS=linux GOARCH=arm go build -o bin/sambal-linux-arm

echo "Done! compiled into bin/"
