# drone-hg

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-hg/status.svg)](http://beta.drone.io/drone-plugins/drone-hg)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-hg/coverage.svg)](https://aircover.co/drone-plugins/drone-hg)
[![](https://badge.imagelayers.io/plugins/drone-hg:latest.svg)](https://imagelayers.io/?images=plugins/drone-hg:latest 'Get your own badge on imagelayers.io')

Drone plugin to clone `mercurial` repositories

## Overview

This plugin is responsible for cloning `mercurial` repositories. It is capable
of cloning a specific commit, branch, tag or pull request. The clone path is
provided in the `dir` field.

## Binary

Build the binary using `make`:

```
make deps build
```

### Clone a commit

```sh
./drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```

### Clone a pull request

```sh
./drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "pull_request",
        "branch": "master",
        "commit": "8d6a233744a5dcacbf2605d4592a4bfe8b37320d",
        "ref": "refs/pull/892/merge"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```

### Clone a tag

```sh
./drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "tag",
        "branch": "master",
        "commit": "339fb92b9629f63c0e88016fffb865e3e1055483",
        "ref": "refs/tags/v0.2.0"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Clone a commit

```sh
docker run -i plugins/drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```

### Clone a pull request

```sh
docker run -i plugins/drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "pull_request",
        "branch": "master",
        "commit": "8d6a233744a5dcacbf2605d4592a4bfe8b37320d",
        "ref": "refs/pull/892/merge"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```

### Clone a tag

```sh
docker run -i plugins/drone-hg <<EOF
{
    "repo": {
        "clone": "https://drone@bitbucket.org/drone/drone"
    },
    "build": {
        # FIXME: How does it look with hg?
        "event": "tag",
        "branch": "master",
        "commit": "339fb92b9629f63c0e88016fffb865e3e1055483",
        "ref": "refs/tags/v0.2.0"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/bitbucket.org/drone/drone"
    }
}
EOF
```
