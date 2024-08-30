FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && apk cache clean
ENV CGO_ENABLED=1
RUN go build -o ./jwt-auth ./cmd/jwt-auth/main.go



FROM alpine
WORKDIR /app
COPY --from=builder /app/jwt-auth ./jwt-auth
COPY --from=builder /app/internal/database/migrations ./internal/database/migrations
EXPOSE 8081
CMD ["./jwt-auth"]