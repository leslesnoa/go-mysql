FROM golang:1.17.4-alpine

COPY . ./root/go

WORKDIR /root/go

RUN apk update && go mod init root/go

WORKDIR /root/go/api

CMD [ "go", "run", "main.go" ]