#!/bin/bash

cd cmd/gaap/

# compilar e comprimir
# linux
go build -ldflags "-s -w" > /dev/null 2>&1

# windows
GOOS=windows GOARCH=386 go build -ldflags "-s -w" > /dev/null 2>&1

mkdir -p build
mv gaap* build/

../../compress.sh