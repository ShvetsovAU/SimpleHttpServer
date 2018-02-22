FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/SimpleHttpServer

RUN go install github.com/shvetsovau/SimpleHttpServer

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/SimpleHttpServer"]

EXPOSE 9000