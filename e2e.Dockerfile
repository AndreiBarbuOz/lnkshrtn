FROM golang:1.16.0 as lnkshrtn-build

WORKDIR /go/src/github.com/AndreiBarbuOz/lnkshrtn

COPY go.mod go.mod

RUN go mod download

# Perform the build
COPY . .
RUN make build_instrumented

# Image

FROM ubuntu:20.10

COPY --from=lnkshrtn-build /go/src/github.com/AndreiBarbuOz/lnkshrtn/dist/main_cover /usr/local/bin/main

USER root
RUN ln -s /usr/local/bin/main /usr/local/bin/lnkshrtn-server
RUN ln -s /usr/local/bin/main /usr/local/bin/lnkshrtn-sidecar

USER 999
