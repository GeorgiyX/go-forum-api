# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd

RUN go build -o /api.run cmd/main.go

# Deploy stage
FROM alpine

WORKDIR /
COPY --from=build /api.run /api.run
EXPOSE 8080

CMD ["/api.run"]
VOLUME ["/vol"]
