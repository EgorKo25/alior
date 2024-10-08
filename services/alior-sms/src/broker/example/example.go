package main

import (
	"alior-sms/src/broker"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"alior-sms/src/logger"

	"github.com/streadway/amqp"
)

func main() {
	ctx, done := context.WithCancel(context.Background())

	writing := make(chan amqp.Publishing)
	reading := make(chan amqp.Delivery)

	logs := logger.NewZapLogger()

	cfg := broker.NewDialConfig("amqp://guest:guest@localhost:5672/", "test-exchange", "direct", "test-key", "testin-queue", true)
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg, logs)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Publish(dial, writing, logs)
		done()
	}()
	go func() {
		dial, err := broker.DialSessionChan(ctx, *cfg, logs)
		if err != nil {
			log.Fatalf(err.Error())
		}

		broker.Subscribe(dial, reading, logs)
		done()
	}()

	// Пример - чтение ввода с консоли и отправка в канал writing
	go func() {
		scanner := bufio.NewScanner(os.Stdin)

		for {
			if scanner.Scan() {
				text := scanner.Text()
				writing <- amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(text),
				}
			}

			if err := scanner.Err(); err != nil {
				log.Printf("Error reading from console: %v", err)
				done()

				return
			}
		}
	}()

	// Пример - чтение сообщений из канала reading и вывод в консоль
	go func() {
		for msg := range reading {
			fmt.Printf("Received message: %s\n", string(msg.Body))
		}
	}()

	<-ctx.Done()
}
