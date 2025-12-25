FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY tools/seed ./tools/seed

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o seed ./tools/seed/cmd/app

FROM alpine:3.20

WORKDIR /app
COPY --from=build /app/seed /app/seed

CMD ["./seed"]
