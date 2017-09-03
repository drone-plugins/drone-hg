# drone-hg

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-hg/status.svg)](http://beta.drone.io/drone-plugins/drone-hg)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-hg?status.svg)](http://godoc.org/github.com/drone-plugins/drone-hg)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-hg)](https://goreportcard.com/report/github.com/drone-plugins/drone-hg)
[![](https://images.microbadger.com/badges/image/plugins/hg.svg)](https://microbadger.com/images/plugins/hg "Get your own image badge on microbadger.com")

Drone plugin to clone `mercurial` repositories. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-hg/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-hg
docker build --rm -t plugins/hg .
```

## Usage

Clone a commit:

```
docker run --rm \
  -e DRONE_REMOTE_URL=https://bitbucket.org/cedk/drone-hg-test \
  -e DRONE_WORKSPACE=/go/src/bitbucket.org/cedk/drone-hg-test \
  -e DRONE_BUILD_EVENT=push \
  -e DRONE_COMMIT_SHA=37526193d0139f188b20e5c8bed8fc0640c38627 \
  plugins/hg
```
