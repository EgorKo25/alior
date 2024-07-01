FROM golang:1.22.4-alpine

LABEL authors="alior-team"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o start.auth main.go

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN chown -R appuser:appgroup /app

USER appuser

# Используем переменные окружения для передачи флага dbhost
CMD ["sh", "-c", "./start.auth -dbhost ${DB_HOST}"]
