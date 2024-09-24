FROM golang:1.22

ENV GONOPROXY=github.com \
    GONOSUMDB=github.com \
    GOPRIVATE=github.com \
    CGO_ENABLED=0

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY go.mod .

COPY .ssh/* /root/.ssh/

RUN chmod 400 /root/.ssh/*

RUN git config --global url."ssh://git@github.com".insteadOf "https://github.com"

RUN mkdir -p /root/.ssh && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN go mod download

COPY . .

RUN go build -o start.alior-cms main.go

CMD ["sh", "-c", "./start.alior-cms"]