# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
COPY src/ ./

RUN go build -o /audit-server

CMD ["/audit-server"]
