FROM golang:1.26-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /build

RUN go build -o /build/shop-service ./cmd/app

RUN go build -o /build/consumer ./cmd/consumer

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /build/shop-service /app/shop-service
COPY --from=builder /build/consumer /app/consumer

COPY configs /app/configs

EXPOSE 8080

CMD ["/app/shop-service"]