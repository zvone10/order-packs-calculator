FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/order-pack-calculator-api ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/bin/order-pack-calculator-api .
EXPOSE 5000

CMD ["./order-pack-calculator-api"]
