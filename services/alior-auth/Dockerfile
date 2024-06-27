FROM golang:latest
LABEL authors="alior-team"

RUN mkdir -p /alior-digital/config
RUN mkdir -p /alior-digital/migration
RUN mkdir -p /alior-digital/src

COPY config/* /alior-digital/config/
COPY migrations/* /alior-digital/migration/
COPY src/* /alior-digital/src/
COPY main.go /alior-digital/
COPY go.mod /alior-digital/

WORKDIR /alior-digital

RUN go mod tidy
RUN go build -o start.auth main.go

CMD ["./start.auth"]
