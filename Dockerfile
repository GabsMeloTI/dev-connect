# Etapa de Build
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init

COPY db/migration /app/db/migration

RUN go build -o main ./main.go

# Etapa final
FROM alpine:3.18

RUN apk add --no-cache ca-certificates curl

WORKDIR /app

COPY --from=builder /app/main /app/main

# Certifique-se de que as migrações sejam copiadas também
COPY --from=builder /app/db/migration /app/db/migration

COPY .env /app/.env

RUN chmod +x /app/main

EXPOSE 8000

CMD ["/app/main"]
