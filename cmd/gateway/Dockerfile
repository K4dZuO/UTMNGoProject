FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o gateway ./cmd/gateway

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/gateway .
COPY cmd/gateway/gateway_config.yaml ./gateway_config.yaml
COPY docs ./docs

EXPOSE 8080


CMD ["./gateway"]
