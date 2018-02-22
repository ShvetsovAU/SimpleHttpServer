FROM golang:latest

ADD . $GOPATH/src/github.com/shvetsovau/SimpleHttpServer
ADD cert.pem /go/bin/cert.pem
ADD key.pem /go/bin/key.pem

RUN go get

RUN go install github.com/shvetsovau/SimpleHttpServer

WORKDIR /go/bin
ENTRYPOINt ["/go/bin/SimpleHttpServer"]

EXPOSE 9000