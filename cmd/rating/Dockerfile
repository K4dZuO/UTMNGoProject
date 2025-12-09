FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o rating ./cmd/rating

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/rating .

ENV PORT=8082

CMD ["./rating"]
