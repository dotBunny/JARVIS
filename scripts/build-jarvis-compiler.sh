#!/bin/bash
cd "$(dirname "$0")"

# Go back a folder
cd ..

# Make sure the 'bin' folder is actually there
mkdir -p bin/jarvis-compiler

# COMPILE!
go build -o bin/jarvis-compiler/jarvis-compiler ./src/jarvis-compiler