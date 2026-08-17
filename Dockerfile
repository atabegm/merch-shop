FROM golang:1.26-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go build -o /build ./internal/cmd \
    && go clean -cache -modcache

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /build /build

EXPOSE 8080

CMD ["/build"]
