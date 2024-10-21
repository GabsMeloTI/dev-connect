FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main

RUN chmod +x main

RUN ls -l main
RUN whoami
RUN id

EXPOSE 8000

CMD ["./main"]