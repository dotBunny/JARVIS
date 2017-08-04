
#!/bin/bash
cd "$(dirname "$0")"

# Remove & Create Temp Folder
rm -rf jarvis-build/
mkdir -p jarvis-build

# Build OSX
mkdir -p jarvis-build/macOS
cp -rf resources/build/macos/JARVIS.app jarvis-build/macOS
go build -o jarvis-build/macOS/JARVIS.app/Contents/MacOS/jarvis ./src
cp -rf resources/www jarvis-build/macOS/JARVIS.app/Contents/Resources/
cp -rf resources/scripts jarvis-build/macOS/JARVIS.app/Contents/Resources/
cp -rf resources/config jarvis-build/macOS/JARVIS.app/Contents/Resources/
./resources/build/tools/configurator/configurator ~/Dropbox/The\ Forgotten\ Soul/Stream/Apps/JARVIS/jarvis.csv  jarvis-build/macOS/JARVIS.app/Contents/Resources/config

# Copy In Override

# Move to Applications
rm -rf /Applications/JARVIS.app
cp -rf jarvis-build/macOS/JARVIS.app /Applications

# Remove Temp Folder
rm -rf jarvis-build/