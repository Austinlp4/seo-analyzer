#!/bin/bash

# Function to stop all background processes on exit
cleanup() {
    echo "Stopping all processes..."
    kill $(jobs -p)
    exit
}

# Set up the cleanup function to run on script exit
trap cleanup EXIT

# Start the Go server
echo "Starting Go server..."
go run backend/cmd/server/main.go &

# Use fswatch for file watching on macOS
echo "Watching frontend files..."
fswatch -o frontend/public frontend/pages | while read f; do
    echo "Changes detected, refreshing..."
done
