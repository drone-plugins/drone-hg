FROM alpine:3.5

RUN apk update && \
    apk add \
        ca-certificates \
        mercurial \
        openssh \
        curl \
        perl && \
    rm -rf /var/cache/apk/*

ADD drone-hg /bin/
ENTRYPOINT ["/bin/drone-hg"]
