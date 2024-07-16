FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY services/cms .

RUN go build -o start.cms ./cmd/cms/main.go

CMD ["sh", "-c", "./start.cms --config=./config/local.yaml"]