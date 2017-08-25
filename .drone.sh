#!/bin/sh

set -e
set -x

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/amd64/drone-hg
