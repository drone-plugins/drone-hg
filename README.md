# drone-hg

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-hg/status.svg)](http://cloud.drone.io/drone-plugins/drone-hg)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/hg.svg)](https://microbadger.com/images/plugins/hg "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-hg?status.svg)](http://godoc.org/github.com/drone-plugins/drone-hg)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-hg)](https://goreportcard.com/report/github.com/drone-plugins/drone-hg)

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
