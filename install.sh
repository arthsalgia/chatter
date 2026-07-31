#!/bin/bash
set -e

REPO="arthsalgia/chatter"
BINARY_NAME="chatter"

if [ "$(uname -s)" != "Darwin" ]; then
    echo "Error: This tool is only supported on macOS."
    exit 1
fi

ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    ZIP_NAME="chatter-darwin-amd64.zip"
elif [ "$ARCH" = "arm64" ]; then
    ZIP_NAME="chatter-darwin-arm64.zip"
else
    echo "Error: Unsupported architecture: $ARCH"
    exit 1
fi

URL="https://github.com/$REPO/releases/latest/download/$ZIP_NAME"

echo "Downloading $BINARY_NAME for macOS ($ARCH)..."
TMP_DIR=$(mktemp -d)
curl -L -o "$TMP_DIR/$ZIP_NAME" "$URL"

echo "Unpacking..."
unzip -q "$TMP_DIR/$ZIP_NAME" -d "$TMP_DIR"

# Find the first file inside the temp directory that isn't the zip file
EXTRACTED_FILE=$(find "$TMP_DIR" -maxdepth 2 -type f ! -name "*.zip" | head -n 1)

if [ -z "$EXTRACTED_FILE" ]; then
    echo "Error: No files found inside the downloaded zip."
    rm -rf "$TMP_DIR"
    exit 1
fi

echo "Installing to /usr/local/bin..."
sudo mv "$EXTRACTED_FILE" /usr/local/bin/$BINARY_NAME
sudo chmod +x /usr/local/bin/$BINARY_NAME

# Clear macOS quarantine flag so it doesn't block the app
sudo xattr -d com.apple.quarantine /usr/local/bin/$BINARY_NAME 2>/dev/null || true

# Cleanup
rm -rf "$TMP_DIR"

echo "Installed successfully!"
echo "To run it, either copy your 'chat.db' into your current folder,"
echo "   or grant your Terminal 'Full Disk Access' in macOS System Settings."
echo "   Then simply type: $BINARY_NAME"