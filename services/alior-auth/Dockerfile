FROM golang:latest
LABEL authors="alior-team"

RUN mkdir -p /app/config
RUN mkdir -p /app/migration
RUN mkdir -p /app/src

COPY config/* /app/config/
COPY migrations/* /app/migration/
COPY src/* /app/src/
COPY main.go /app/
COPY go.mod /app/

WORKDIR /app

RUN go mod tidy
RUN go build -o start.auth main.go

CMD ["./start.auth"]
