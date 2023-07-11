# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /mapDemo

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /mapDemo /mapDemo

EXPOSE 30022

ENTRYPOINT ["/mapDemo"]