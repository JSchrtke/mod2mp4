#!/bin/bash
rm -r build
mkdir build

# build for linux
mkdir ./build/mod2mp4-linux64 
env GOOS=linux GOARCH=amd64 go build -o ./build/mod2mp4-linux64/mod2mp4 .
mkdir ./build/mod2mp4-linux64/deps
cp -r ./deps/ffmpeg-linux64/ ./build/mod2mp4-linux64/deps/ffmpeg-linux64/

# build for windows
mkdir build/mod2mp4-win64 
env GOOS=windows GOARCH=amd64 go build -o build/mod2mp4-win64/mod2mp4.exe -ldflags -H=windowsgui . 
mkdir ./build/mod2mp4-win64/deps
cp -r ./deps/ffmpeg-win64/ ./build/mod2mp4-win64/deps/ffmpeg-win64/

# pack everything up
7z a build/mod2mp4-linux64.zip build/mod2mp4-linux64/*
7z a build/mod2mp4-win64.zip build/mod2mp4-win64/*
