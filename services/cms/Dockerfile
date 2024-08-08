FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o start.cms main.go

ARG CONFIG_FILE=config.yaml
ENV CONFIG_FILE=${CONFIG_FILE}

CMD ["sh", "-c", "./start.cms --config=./config/${CONFIG_FILE}"]