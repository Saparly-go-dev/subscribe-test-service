FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY configs/config.yaml ./configs/config.yaml
COPY .env .env  

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/.env ./.env


EXPOSE 8085

CMD ["./server"]
