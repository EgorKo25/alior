package main

import (
	"alior-sms/src/broker"
	"context"
	"log"
	"os"
)

func main() {
	ctx, done := context.WithCancel(context.Background())

	cfg := broker.NewDialConfig("amqp://guest:guest@localhost:5672/", "test-exchange", "direct", "test-key", "testin-queue", true)
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Publish(dial, broker.MakeReaderChan(os.Stdin))
		done()
	}()
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Subscribe(dial, broker.MakeWriterChan(os.Stdout))
		done()
	}()
	<-ctx.Done()
}
