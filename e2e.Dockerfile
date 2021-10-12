FROM golang:1.16.0 as lnkshrtn-build

WORKDIR /go/src/github.com/AndreiBarbuOz/lnkshrtn

COPY go.mod go.mod

RUN go mod download

# Perform the build
COPY . .
RUN make build_instrumented

# Image

FROM ubuntu:20.10

COPY --from=lnkshrtn-build /go/src/github.com/AndreiBarbuOz/lnkshrtn/dist/lnk-server-dbg /usr/local/bin/lnk-server

USER 999
