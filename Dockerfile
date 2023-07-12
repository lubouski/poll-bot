FROM golang:1.20.3-alpine3.17 as dev

WORKDIR /work

RUN apk add vim git openssh
