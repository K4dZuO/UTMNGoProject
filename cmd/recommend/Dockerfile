FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o recommend ./cmd/recommend

FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/recommend .
COPY cmd/recommend/recommend_config.yaml ./recommend_config.yaml

EXPOSE 8083

CMD ["./recommend"]
