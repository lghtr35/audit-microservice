# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
RUN apk add build-base

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
COPY src/ ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN go build -o /audit-server

CMD ["/audit-server"]
