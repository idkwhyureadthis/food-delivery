FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./food-service ./cmd/food-service/main.go



FROM alpine
WORKDIR /app
COPY --from=builder /app/food-service ./food-service
EXPOSE 8080
CMD ["./food-service"]