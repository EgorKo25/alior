FROM golang:1.22-alpine

RUN apk add --no-cache git openssh-client

WORKDIR /app

COPY .ssh/egorko25_private /root/.ssh/egorko25_private
COPY .ssh/config /root/.ssh/config

RUN chmod 600 /root/.ssh/egorko25_private && \
    ssh-keyscan -t rsa github.com >> /root/.ssh/known_hosts

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o start.alior-cms main.go

CMD ["sh", "-c", "./start.alior-cms"]