FROM golang:1.16-alpine as base

RUN GOCACHE=OFF

WORKDIR /src/app

COPY go.mod go.sum /src/app/
RUN go mod download

COPY app /src/app/app
COPY gen /src/app/gen
COPY migrations /src/app/migrations

RUN go build -o /go/bin/meta-server-app ./gen/swagger/cmd/meta-server

FROM alpine
RUN mkdir -p /go/bin
WORKDIR /go/bin
COPY --from=base /go/bin/meta-server-app .
COPY --from=base /src/app/migrations ./migrations

ENTRYPOINT ["/go/bin/meta-server-app"]