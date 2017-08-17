#!/bin/bash
cd "$(dirname "$0")"

cd ..

# Remove & Create Temp Folder
rm -rf jarvis-build/
mkdir -p jarvis-build

# Build OSX
mkdir -p jarvis-build/macOS

cp -rf resources/templates/macos/JARVIS.app jarvis-build/macOS
go build -o jarvis-build/macOS/JARVIS.app/Contents/MacOS/jarvis ./src/jarvis

# Copy OSX Resources
cp -rf resources/www jarvis-build/macOS/JARVIS.app/Contents/Resources/
cp -rf resources/scripts jarvis-build/macOS/JARVIS.app/Contents/Resources/
cp -rf resources/config jarvis-build/macOS/JARVIS.app/Contents/Resources/

# Build For Windows
mkdir -p jarvis-build/windows
rsrc -manifest resources/templates/windows/jarvis.exe.manifest -ico resources/templates/windows/jarvis.ico -o src/jarvis/jarvis.syso
GOOS=windows GOARCH=386 go build -o jarvis-build/windows/jarvis.exe ./src/jarvis
rm -rf src/jarvis/jarvis.syso
cp -rf resources/www jarvis-build/windows/
cp -rf resources/scripts jarvis-build/windows/
cp -rf resources/config jarvis-build/windows/

# Remove Previous Package
rm -rf jarvis-build.zip

# Copy Over Instructions
cp -rf README.md jarvis-build/

# Compress File
zip -r jarvis-build.zip jarvis-build/

# Remove Temp Folder (if there are no arguements)
if [ $# -eq 0 ]
then
    rm -rf jarvis-build/
fi