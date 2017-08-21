
#!/bin/bash

# TODO: Make these stop if there is an error

# Define with wildcard because of space
CONFIG="/Users/reapazor/Dropbox/The Forgotten Soul/Stream/Apps/JARVIS/jarvis.csv"

cd "$(dirname "$0")"
clear

if [ $# -eq 0 ]
then
    echo "Building Compiler ..."
    ./scripts/build-jarvis-compiler.sh
    echo "Building JARVIS ..."
    ./scripts/build-jarvis.sh "$CONFIG"
    echo "Running JARVIS ..."
    cd bin/jarvis
    ./jarvis
else
    echo "Building Compiler ..."
    ./scripts/build-jarvis-compiler.sh

    echo "Building Package ..."
    ./scripts/package.sh DONOTDELETE

    echo "Setting Configuration"
    ./bin/jarvis-compiler/jarvis-compiler "$CONFIG" "./jarvis-build/macOS/JARVIS.app/Contents/Resources/config"

    echo "Copying Into Applications..."
    rm -rf /Applications/JARVIS.app
    cp -rf jarvis-build/macOS/JARVIS.app /Applications

    echo "Clearing Build Files..."
    rm -rf jarvis-build/
    rm -rf jarvis-build.zip
fi