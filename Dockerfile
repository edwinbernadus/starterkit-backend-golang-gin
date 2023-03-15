# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./

COPY models/ /app/models/
COPY module_socket/ /app/module_socket/
COPY module_db/ /app/module_db/

RUN go mod download

RUN go build -o one

EXPOSE 8080

ENTRYPOINT ["./one"]