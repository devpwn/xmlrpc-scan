#!/bin/bash

# Name of your Go program
PROGRAM_NAME="xmlrpcscan"

# Build for Linux (AMD64)
echo "Building for Linux (AMD64)..."
env GOOS=linux GOARCH=amd64 go build -o "${PROGRAM_NAME}_linux_amd64" ../xmlrpcscan.go

# Build for Linux (ARM64)
echo "Building for Linux (ARM64)..."
env GOOS=linux GOARCH=arm64 go build -o "${PROGRAM_NAME}_linux_arm64" ../xmlrpcscan.go

# Build for macOS (AMD64)
echo "Building for macOS (AMD64)..."
env GOOS=darwin GOARCH=amd64 go build -o "${PROGRAM_NAME}_darwin_amd64" ../xmlrpcscan.go

# Build for macOS (ARM64)
echo "Building for macOS (ARM64)..."
env GOOS=darwin GOARCH=arm64 go build -o "${PROGRAM_NAME}_darwin_arm64" ../xmlrpcscan.go

# Build for Windows (AMD64)
echo "Building for Windows (AMD64)..."
env GOOS=windows GOARCH=amd64 go build -o "${PROGRAM_NAME}_windows_amd64.exe" ../xmlrpcscan.go

# Build for Windows (ARM64)
echo "Building for Windows (ARM64)..."
env GOOS=windows GOARCH=arm64 go build -o "${PROGRAM_NAME}_windows_arm64.exe" ../xmlrpcscan.go

echo "Build process completed successfully!"
