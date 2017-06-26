
#!/bin/bash
cd "$(dirname "$0")"

# Remove & Create Temp Folder
rm -rf jarvis-build/
mkdir -p jarvis-build

# Build OSX
mkdir -p jarvis-build/macOS
cp -rf resources/build/macos/JARVIS.app jarvis-build/macOS
cp -rf bin/jarvis.json jarvis-build/macOS/JARVIS.app/Contents/Resources/
go build -o jarvis-build/macOS/JARVIS.app/Contents/MacOS/jarvis ./src
cp -rf resources/www jarvis-build/macOS/JARVIS.app/Contents/Resources/

# Move to Applications
rm -rf /Applications/JARVIS.app
cp -rf jarvis-build/macOS/JARVIS.app /Applications

# Remove Temp Folder
rm -rf jarvis-build/