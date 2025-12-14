FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (required for go get)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o rating_service cmd/rating/main.go

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /app/rating_service .
COPY cmd/rating/rate_config.yaml ./rate_config.yaml

EXPOSE 8082

CMD ["./rating_service"]
