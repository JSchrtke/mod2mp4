#!/bin/bash
rm -r build
mkdir build
pushd build

# build for linux
mkdir ./mod2mp4-linux64
env GOOS=linux GOARCH=amd64 go build -o ./mod2mp4-linux64/mod2mp4 ..
mkdir ./mod2mp4-linux64/deps
cp -r ../deps/ffmpeg-linux64/ ./mod2mp4-linux64/deps/ffmpeg-linux64/

# build for windows
env GOOS=windows GOARCH=amd64 go build -o ./mod2mp4-win64/mod2mp4.exe -ldflags -H=windowsgui ..
mkdir ./mod2mp4-win64/deps
cp -r ../deps/ffmpeg-win64/ ./mod2mp4-win64/deps/ffmpeg-win64/

# pack everything up
7z a ./mod2mp4-linux64.zip ./mod2mp4-linux64/*
rm ./mod2mp4-win64.zip
7z a ./mod2mp4-win64.zip ./mod2mp4-win64/*
