#!/bin/bash
cd "$(dirname "$0")"
clear

rm -rf ./bin
mkdir ./bin

# Make For macOS
mkdir -p temp/macOS
cp -rf ../../macos/JARVISHelper.app temp/macOS
go build -o temp/macOS/JARVISHelper.app/Contents/MacOS/jarvisHelper ./src
cp -r temp/macOS/JARVISHelper.app ./bin
rm -rf temp/