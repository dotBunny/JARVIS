
#!/bin/bash
cd "$(dirname "$0")"
go build -o bin/jarvis ./src
GOOS=windows GOARCH=386 go build -o bin/jarvis.exe ./src
rm -rf deploy/
rm -rf jarvis.zip
mkdir -p deploy
cp -rf bin/jarvis deploy/
cp -rf bin/jarvis.exe deploy/
cp -rf resources/jarvis.toml deploy/
zip -r jarvis.zip deploy/
rm -rf deploy/