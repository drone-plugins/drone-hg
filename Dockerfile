FROM plugins/base:amd64
MAINTAINER Drone.IO Community <drone-dev@googlegroups.com>

RUN apk add -U --no-cache mercurial openssh curl perl

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/drone-plugins/drone-hg.git"
LABEL org.label-schema.name="Drone Mercurial"
LABEL org.label-schema.vendor="Drone.IO Community"
LABEL org.label-schema.schema-version="1.0"

ADD release/linux/amd64/drone-hg /bin/
ENTRYPOINT [ "/bin/drone-hg" ]
