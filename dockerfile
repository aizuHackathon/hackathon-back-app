FROM golang:1.18-alpine

RUN apk update &&  apk add git

WORKDIR /go/src/app

ADD ./go .

RUN go mod tidy