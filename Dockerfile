FROM golang:1.21.4-alpine as builder

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -o /usr/local/bin/dialog-service cmd/dialog/main.go

EXPOSE 8070 8071 7050