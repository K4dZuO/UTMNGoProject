FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o reviews ./cmd/reviews

FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/reviews .
COPY cmd/reviews/config.yaml ./config.yaml
COPY --from=build /app/migrations ./migrations

EXPOSE 8081

CMD ["./reviews"]
