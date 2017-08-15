# drone-hg

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-hg/status.svg)](http://beta.drone.io/drone-plugins/drone-hg)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-hg/coverage.svg)](https://aircover.co/drone-plugins/drone-hg)
[![](https://badge.imagelayers.io/plugins/drone-hg:latest.svg)](https://imagelayers.io/?images=plugins/drone-hg:latest 'Get your own badge on imagelayers.io')

Drone plugin to clone `mercurial` repositories. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

This plugin is responsible for cloning `mercurial` repositories. It is capable
of cloning a specific commit, branch, tag or pull request. The clone path is
provided in the `dir` field.

## Build

Build the binary with the following commands:

```
go build
go test
```
## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/hg .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-hg' not found or does not exist..
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
