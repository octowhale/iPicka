#!/bin/bash
#

source ./version
case $1 in 
linux)  CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -o ipicka-linux-${VERSION} ;;
darwin) CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ipicka-darwin-${VERSION} ;;
*)      go build -o ipicka ;;
esac