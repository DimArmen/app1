# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY *.go ./

RUN go env -w GO111MODULE=off
RUN go build -o /goserver

CMD ["/goserver"]