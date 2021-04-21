FROM golang:1.15-alpine

COPY . /src/app

WORKDIR /src/app

RUN env GOOS=linux GOARCH=amd64 go build -o ./dist/meta-server ./gen/swagger/cmd/meta-server
