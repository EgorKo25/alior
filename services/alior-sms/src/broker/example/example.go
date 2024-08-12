package main

import (
	"alior-sms/src/broker"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func read(r io.Reader) <-chan []byte {
	lines := make(chan []byte)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			lines <- []byte(scan.Bytes())
		}
	}()
	return lines
}

func write(w io.Writer) chan<- []byte {
	lines := make(chan []byte)
	go func() {
		for line := range lines {
			fmt.Fprintln(w, string(line))
		}
	}()
	return lines
}

func main() {
	ctx, done := context.WithCancel(context.Background())

	cfg := broker.NewDialConfig("amqp://guest:guest@localhost:5672/", "test-exchange", "direct", "test-key", "testin-queue", true)
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg)
		if err != nil {
			log.Fatalf(err.Error())
		}
		broker.Publish(dial, read(os.Stdin))
		done()
	}()
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg)
		if err != nil {
			log.Fatalf(err.Error())
		}
		broker.Subscribe(dial, write(os.Stdout))
		done()
	}()
	<-ctx.Done()
}
