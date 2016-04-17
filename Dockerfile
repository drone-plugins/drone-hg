# Docker image for Drone's hg-clone plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-hg .

FROM alpine:3.2
RUN apk add -U ca-certificates mercurial openssh curl perl && rm -rf /var/cache/apk/*
ADD drone-hg /bin/
ENTRYPOINT ["/bin/drone-hg"]
