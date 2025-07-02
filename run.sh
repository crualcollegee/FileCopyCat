#!/bin/bash

export CGO_ENABLED=0

# Check if "github.com/spf13/cobra" is installed
if ! go list -m github.com/spf13/cobra >/dev/null 2>&1; then
    echo "Dependency 'github.com/spf13/cobra' not installed, installing..."
    go get github.com/spf13/cobra
else
    echo "'github.com/spf13/cobra' is already installed"
fi

# Compile the Go project
echo "Compiling Go project..."
go build -o main ./main.go
if [ $? -ne 0 ]; then
    echo "Build failed. Please check the code."
    exit 1
fi
echo "Build successful!"

# Interactive command execution
while true; do
    # Prompt user for command input
    echo "Available command:"
    echo " ExtCopy [sourceDir] [targetDir] [extension]"
    echo " ExtMove [sourceDir] [targetDir] [extension]"
    echo "or type 'help' for instructions, 'exit' to quit:"
    read -p "> " command
    
    if [ "$command" == "exit" ]; then
        echo "Exiting the script."
        break
    fi
    
    # Execute the command
    ./main $command

    # Check the result of the command
    if [ $? -eq 0 ]; then
        echo "Command executed successfully."
    else
        echo "Command execution failed."
    fi
done
