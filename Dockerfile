FROM golang:1.22.0

WORKDIR /app


COPY . /app

RUN go mod tidy

RUN  go run main.go -b 0.0.0.0