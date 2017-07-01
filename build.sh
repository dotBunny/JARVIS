#!/bin/bash
cd "$(dirname "$0")"
clear
mkdir -p bin
go build -o bin/jarvis ./src
rm -rf bin/www
mkdir -p bin/www
cp -rf resources/www bin/
cp -rf resources/scripts bin/
