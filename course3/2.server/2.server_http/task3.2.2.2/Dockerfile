# syntax=docker/dockerfile:1

FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o task3222 .

CMD ["./task3222"]