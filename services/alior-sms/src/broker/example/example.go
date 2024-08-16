package main

import (
	"alior-sms/src/broker"
	"context"
	"log"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func main() {
	ctx, done := context.WithCancel(context.Background())

	var (
		writing chan amqp.Publishing
		reading chan amqp.Delivery
	)

	logger := zap.Logger{}

	cfg := broker.NewDialConfig("amqp://guest:guest@localhost:5672/", "test-exchange", "direct", "test-key", "testin-queue", true)
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg, logger)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Publish(dial, writing, logger)
		done()
	}()
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg, logger)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Subscribe(dial, reading, logger)
		done()
	}()

	<-ctx.Done()
}
