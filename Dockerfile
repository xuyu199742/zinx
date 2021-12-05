FROM golang:alpine

LABEL MAINTAINER="Bruce"

WORKDIR /go/src/

COPY . .

EXPOSE 8888