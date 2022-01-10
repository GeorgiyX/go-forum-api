# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get -u github.com/mailru/easyjson/...
ENV PATH=$GOPATH/bin:$PATH

COPY . .
RUN easyjson -all -pkg app/models #TODO: go generate
RUN go build -o /api.run cmd/main.go

# Deploy stage
FROM alpine

ENV PG_VERSION 12

RUN apt -y update && \
    apt install -y postgresql-$PG_VERSION

ENV PG_DEFAULT_USER postgres
ENV PG_FORUM_USER forum_user
ENV PG_PASSWORD forum_user_password
ENV PG_DB_NAME forum
ENV PG_PORT 5432
ENV API_PORT 5000

# Мб проблемы с sudo c systemctl
USER $PG_DEFAULT_USER

RUN systemctl start postgresql && \
    psql --command "CREATE USER $PG_FORUM_USER WITH SUPERUSER PASSWORD '$PG_PASSWORD';" && \
    createdb --owner=$PG_FORUM_USER $PG_DB_NAME && \
    systemctl stop postgresql

ENV ARTIFACT api.run
WORKDIR /
COPY --from=build /$ARTIFACT /$ARTIFACT

#TODO: docker-compose
CMD service postgresql start && \
    psql -h localhost -p $PG_PORT -d $PG_DB_NAME -U $PG_FORUM_USER -q -f ./db/db.sql \
    && ./$ARTIFACT

VOLUME ["/var/lib/postgresql/data"]
EXPOSE $API_PORT