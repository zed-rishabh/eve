FROM golang:1.8.3-alpine

RUN apk add --no-cache ca-certificates gcc libc6-compat gnupg make automake wget git coreutils cpio squashfs-tools autoconf file xz patch bc acl-dev openssl-dev at-spi2-core python pkgconfig zlib-dev musl-dev pixman libcap libfdt bash

RUN git clone https://github.com/mgoltzsche/rkt-compose.git
WORKDIR /go/rkt-compose
RUN git reset --hard tags/v0.1.0
COPY *.patch ./
RUN for p in *.patch ; do patch -p1 < "$p" ; done

RUN ./make.sh
