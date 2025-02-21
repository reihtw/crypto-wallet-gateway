FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /crypto-wallet-gateway

EXPOSE 8080

CMD ["/crypto-wallet-gateway"]
