#!/bin/bash
env GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui . && env GOOS=linux GOARCH=amd64 go build .
