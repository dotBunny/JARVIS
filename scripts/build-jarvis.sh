#!/bin/bash
cd "$(dirname "$0")"

# Check for arguement
if [ $# -eq 0 ]
  then
    echo "ERROR: A path to a JSON based replacer configuration file is required."
    exit
fi

# Go back a folder
cd ..

# Make sure the 'bin' folder is actually there
mkdir -p bin/jarvis

# COMPILE!
go build -o bin/jarvis/jarvis ./src/jarvis

# Handle WWW Folder
rm -rf bin/jarvis/www
mkdir -p bin/jarvis/www
cp -rf resources/www bin/jarvis

# Handle Config Folder
rm -rf bin/jarvis/config
mkdir -p bin/jarvis/config
cp -rf resources/config bin/jarvis

# Handle Scripts Folder
rm -rf bin/jarvis/scripts
mkdir -p bin/jarvis/scripts
cp -rf resources/scripts bin/jarvis

# Handle Database
rm -rf bin/jarvis/db.new.sqlite
cp -rf resources/database/db.new.sqlite bin/jarvis

# Compile Configs
./bin/jarvis-compiler/jarvis-compiler "$1" ./bin/jarvis/config