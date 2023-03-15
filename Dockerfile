# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./
# COPY /models/ ./
# COPY /module_db/ ./
# COPY /module_socket/ ./
# ADD models /app
# ADD module_socket /app
# ADD module_db /app

# COPY . /

# COPY *.* ./
COPY models/ /app/models/
COPY module_socket/ /app/module_socket/
COPY module_db/ /app/module_db/

RUN go mod download

# RUN go build -o one

EXPOSE 8080

# CMD [ "go run ." ]
ENTRYPOINT ["tail", "-f", "/dev/null"]
# ENTRYPOINT ["one"]