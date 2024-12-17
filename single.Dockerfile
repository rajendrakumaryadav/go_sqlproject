FROM golang:latest AS builder
WORKDIR /app/source
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum /app/source/
RUN go mod download
COPY . /app/source

RUN go build -o webapp .
ENTRYPOINT ["/app/source/webapp"]
