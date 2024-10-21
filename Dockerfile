# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

# Copia o código-fonte e compila
COPY . ./
RUN go build -o main

# Production stage
FROM alpine:3.18

# Cria uma pasta para o binário
WORKDIR /app

# Copia o binário compilado da fase anterior
COPY --from=builder /app/main /app/main

# Permissões para o binário
RUN chmod +x /app/main

# Exposição da porta usada pela aplicação
EXPOSE 8000

# Comando para executar a aplicação
CMD ["/app/main"]
