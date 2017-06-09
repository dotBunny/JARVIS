#!/bin/bash
cd "$(dirname "$0")"
clear
go build -o bin/jarvis ./src
cd bin
./jarvis
