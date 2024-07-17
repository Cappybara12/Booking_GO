#!/bin/bash

# Navigate to the directory containing main.go
cd cmd/web || { echo "Failed to navigate to cmd/web. Exiting."; exit 1; }

# Print the current working directory
echo "Current working directory: $(pwd)"

# Print a message indicating the start of the build process
echo "Building the Go project..."

# Build the Go project, specifying the output path for the binary
go build -o ../../myproject

# Check if the build was successful
if [ $? -eq 0 ]; then
    # Navigate back to the root directory
    cd ../../ || { echo "Failed to navigate back to the root directory. Exiting."; exit 1; }

    # Print the current working directory
    echo "Current working directory: $(pwd)"

    # Print a message indicating the build was successful
    echo "Build successful, running the project..."

    # Run the generated executable
    ./myproject
else
    # Print a message indicating the build failed
    echo "Build failed, please check for errors."
fi
