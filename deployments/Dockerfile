
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o climate-data-service ./cmd/climate-data-service/main.go


FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/climate-data-service .

EXPOSE 50051

ENV CLIMATE_LOG_LEVEL=info

ENTRYPOINT ["./climate-data-service"]
