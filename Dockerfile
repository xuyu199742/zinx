FROM golang:alpine as builder

LABEL MAINTAINER="Bruce"

RUN go build -o server .

FROM alpine:latest

WORKDIR /go/src/zinx

COPY . .

EXPOSE 8888
