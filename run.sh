#!/bin/bash
cd "$(dirname "$0")"
clear
mkdir -p bin
go build -o bin/jarvis ./src
cd bin
./jarvis
