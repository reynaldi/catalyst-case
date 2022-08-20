FROM golang:1.19-alpine AS builder
ENV CONNECTION_STRING=root:zxasqw12@tcp(db:3306)/catalyst_db?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true
ENV DIALECT=mysql
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . .

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin

ENTRYPOINT [ "/app/bin" ]