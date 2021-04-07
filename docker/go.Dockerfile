FROM golang:1.16-alpine

WORKDIR /go/src/app

RUN apk add build-base
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080 2345
CMD ["dlv", "debug", "/go/src/app", "--headless", "--listen=:2345", "--api-version=2", "--log"]