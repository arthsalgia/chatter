#!/bin/bash

clear
echo "========================================="
echo "   iMessage Analyzer Server Setup        "
echo "========================================="
echo ""

LOCAL_FILE="./chat.db"
SYSTEM_SOURCE="$HOME/Library/Messages/chat.db"
DB_DEST="./local_chat.db"

PORT_NUMBER="${1:-8000}"

if ! [[ "$PORT_NUMBER" =~ ^[0-9]+$ ]]; then
    echo "❌ ERROR: '$PORT_NUMBER' is not a valid port number."
    echo "Usage: ./setup.sh [port_number] (e.g., ./setup.sh 8080)"
    exit 1
fi

if [ -f "$LOCAL_FILE" ]; then
    echo "Found local file at $LOCAL_FILE. Prioritizing local data..."
    
    if command -v sqlite3 &> /dev/null; then
        sqlite3 "$LOCAL_FILE" ".backup '$DB_DEST'"
    else
        cp "$LOCAL_FILE" "$DB_DEST"
    fi
    echo "✅ Local database mirrored successfully!"

else
    echo "No local file found. Accessing live system iMessage database..."
    echo "Checking system permissions..."

    head -c 10 "$SYSTEM_SOURCE" > /dev/null 2>&1

    if [ $? -ne 0 ]; then
        echo "❌ ERROR: Operation not permitted."
        echo "------------------------------------------------------------------"
        echo "macOS requires you to grant Full Disk Access to your Terminal app"
        echo "before this tool can securely analyze your local iMessage files."
        echo ""
        echo "Alternative: Place a copy of your 'chat.db' file directly into this"
        echo "project folder, and this script will use it automatically."
        echo "------------------------------------------------------------------"
        exit 1
    fi

    echo "✅ Permissions verified!"
    echo "Syncing system iMessage database safely..."

    if command -v sqlite3 &> /dev/null; then
        sqlite3 "$SYSTEM_SOURCE" ".backup '$DB_DEST'"
    else
        cp "$SYSTEM_SOURCE" "$DB_DEST"
    fi
    echo "✅ System database mirrored successfully!"
fi

echo "Starting your Go API server..."
echo "========================================="
echo ""


if [ -f "./server" ]; then
    ./server -db="$DB_DEST" -port=":$PORT_NUMBER"
else
    go run ./cmd/server/main.go -db="$DB_DEST" -port=":$PORT_NUMBER"
fi

echo "API server successfully started on port :$PORT_NUMBER"
echo "========================================="
echo ""