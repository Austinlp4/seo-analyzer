#!/bin/bash

# Build the React frontend
cd frontend
npm run build
cd ..

# Build the Go backend
go mod tidy
go build -o automated-seo-analyzer backend/cmd/server/main.go
if [ $? -ne 0 ]; then
    echo "Go build failed"
    exit 1
fi

echo "Build completed successfully."