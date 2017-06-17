#!/bin/bash
cd "$(dirname "$0")"
clear
mkdir -p bin
go build -o bin/jarvis ./src
rm -rf bin/resources
mkdir -p bin/resources
cp -rf resources/overlay bin/resources
