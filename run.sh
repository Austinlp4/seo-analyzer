#!/bin/bash

# Build the application
./build.sh

# Check if the binary exists
if [ ! -f "./seo-analyzer" ]; then
    echo "Binary not found. Building..."
    go build -o seo-analyzer backend/cmd/server/main.go
    if [ $? -ne 0 ]; then
        echo "Go build failed"
        exit 1
    fi
fi

# Run the Go server
./seo-analyzer