
#!/bin/bash
cd "$(dirname "$0")"

# Remove & Create Temp Folder
rm -rf jarvis-build/
mkdir -p jarvis-build
mkdir -p jarvis-build/www

# Build OSX
go build -o jarvis-build/jarvis ./src

# Build For Windows
rsrc -manifest resources/build/windows/jarvis.exe.manifest -ico resources/build/windows/jarvis.ico -o src/jarvis.syso
GOOS=windows GOARCH=386 go build -o jarvis-build/jarvis.exe ./src
rm -rf src/jarvis.syso

# Remove Previous Package
rm -rf jarvis.zip

# Copy Resources Over
cp -rf resources/jarvis.json jarvis-build/
cp -rf resources/www jarvis-build/

# Copy Over Instructions
cp -rf README.md jarvis-build/

# Compress File
zip -r jarvis-build.zip jarvis-build/

# Remove Temp Folder
rm -rf jarvis-build/