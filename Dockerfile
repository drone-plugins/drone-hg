FROM alpine:3.5

RUN apk update && \
    apk add \
        ca-certificates \
        mercurial \
        openssh \
        curl \
        perl && \
    rm -rf /var/cache/apk/*

ADD release/linux/amd64/drone-hg /bin/
ENTRYPOINT ["/bin/drone-hg"]
