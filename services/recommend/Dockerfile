FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o recommend ./cmd/app

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/recommend /app/recommend
COPY config.yaml /app/config.yaml

EXPOSE 8083

CMD ["./recommend"]
