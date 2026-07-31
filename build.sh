#!/bin/bash
set -e

echo "📦 Building React frontend with Vite..."
cd chatter-app
npm install
npm run build
cd ..

echo "Compiling Go binary for macOS..."
mkdir -p dist

# Build for Mac Intel
GOOS=darwin GOARCH=amd64 go build -o dist/chatter-darwin-amd64 main.go

# Build for Mac Apple Silicon (M1/M2/M3/M4)
GOOS=darwin GOARCH=arm64 go build -o dist/chatter-darwin-arm64 main.go

echo "Build complete! Binaries are in the dist/ folder."