#!/bin/bash
cd "$(dirname "$0")"
mkdir -p bin
go build -o bin/jarvis ./src