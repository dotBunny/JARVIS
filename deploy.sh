
#!/bin/bash
cd "$(dirname "$0")"

# Remove & Create Temp Folder
rm -rf jarvis-build/
mkdir -p jarvis-build

# Build OSX
mkdir -p jarvis-build/macOS
cp -rf resources/build/macos/JARVIS.app jarvis-build/macOS
cp -rf resources/jarvis.json jarvis-build/macOS/JARVIS.app/Contents/Resources/
go build -o jarvis-build/macOS/JARVIS.app/Contents/MacOS/jarvis ./src
cp -rf resources/www jarvis-build/macOS/JARVIS.app/Contents/Resources/
cp -rf resources/scripts jarvis-build/macOS/JARVIS.app/Contents/Resources/

# Build For Windows
mkdir -p jarvis-build/windows
rsrc -manifest resources/build/windows/jarvis.exe.manifest -ico resources/build/windows/jarvis.ico -o src/jarvis.syso
GOOS=windows GOARCH=386 go build -o jarvis-build/windows/jarvis.exe ./src
rm -rf src/jarvis.syso
cp -rf resources/jarvis.json jarvis-build/windows/
cp -rf resources/www jarvis-build/windows/

# Remove Previous Package
rm -rf jarvis-build.zip

# Copy Over Instructions
cp -rf README.md jarvis-build/

# Compress File
zip -r jarvis-build.zip jarvis-build/

# Remove Temp Folder
rm -rf jarvis-build/