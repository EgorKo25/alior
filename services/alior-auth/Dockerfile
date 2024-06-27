FROM golang:latest
LABEL authors="kozlov-e"

RUN mkdir config
COPY config/* config/
COPY main.go .
COPY go.mod .
RUN go mod download
RUN go build -o start.auth main.go
CMD ["start.auth"]