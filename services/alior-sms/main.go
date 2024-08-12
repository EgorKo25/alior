package main

import (
	"alior-sms/src/broker"
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
)

// read is this application's translation to the message format, scanning from
// stdin.
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

// write is this application's subscriber of application messages, printing to
// stdout.
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
	/*cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}*/

	ctx, done := context.WithCancel(context.Background())

	uri := "amqp://guest:guest@:asd/"
	go func() {
		broker.Publish(broker.Redial(ctx, uri, "Notify", "fanout", "", "Notify_queue"), read(os.Stdin))
		done()
	}()

	go func() {
		broker.Subscribe(broker.Redial(ctx, uri, "Notify", "fanout", "", "Notify_queue"), write(os.Stdout))
		done()
	}()
	<-ctx.Done()
}
