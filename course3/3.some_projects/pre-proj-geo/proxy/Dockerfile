FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o proxy ./cmd/app

#EXPOSE 8080

CMD [ "./proxy" ]
