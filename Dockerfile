FROM plugins/base:amd64

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Mercurial" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache mercurial openssh curl perl

ADD release/linux/amd64/drone-hg /bin/
ENTRYPOINT ["/bin/drone-hg"]
