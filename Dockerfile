FROM golang:1.19-alpine AS builder
ENV CONNECTION_STRING=root:root@tcp(host.docker.internal:3307)/catalyst_db?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true
ENV APP_PORT=4000
ENV APP_HOST=0.0.0.0
ENV DIALECT=mysql
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

EXPOSE 4000

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin

ENTRYPOINT [ "/app/bin" ]