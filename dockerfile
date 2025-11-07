# Dockerfile
FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/climate_data_service ./

FROM gcr.io/distroless/base-debian12
COPY --from=build /bin/climate_data_service /climate_data_service
EXPOSE 8080
ENTRYPOINT ["/climate_data_service"]
