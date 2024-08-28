FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o courier-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/courier-service .

CMD ["./courier-service"]
