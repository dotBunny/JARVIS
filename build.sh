#!/bin/bash
cd "$(dirname "$0")"
mkdir -p bin
# COMPILE!
go build -o bin/jarvis ./src
# Handle WWW Folder
rm -rf bin/www
mkdir -p bin/www
cp -rf resources/www bin/
# Handle Config Folder
rm -rf bin/config
mkdir -p bin/config
cp -rf resources/config bin/
# Handle Scripts Folder
rm -rf bin/scripts
mkdir -p bin/scripts
cp -rf resources/scripts bin/
./resources/build/tools/configurator/configurator ~/Dropbox/The\ Forgotten\ Soul/Stream/Apps/JARVIS/jarvis.csv  ./bin/config