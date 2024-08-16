package serviceWorker

import (
	"alior-sms/src/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const (
	workerCount int = 4
)

type DB interface {
	InsertService(ctx context.Context, service *types.Service) (int32, error)
	GetPaginatedServices(ctx context.Context, limit, offset int32) ([]*types.Service, error)
}

type Job interface {
	DoJob(context.Context, DB) ([]byte, error)
}

func parseMessage(message []byte) (Job, error) {
	var amqpMessage types.AMQPMessage
	err := json.Unmarshal(message, &amqpMessage)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}

	switch amqpMessage.Properties.Type {
	case "create":
		var service types.Service
		err := json.Unmarshal([]byte(amqpMessage.Body), &service)

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal service body: %w", err)
		}

		return InsertJob{service: &service}, nil

	default:
		return nil, fmt.Errorf("unknown content-type: %s", amqpMessage.Properties.ContentType)
	}
}

func Worker(ctx context.Context, db DB, id int, messages <-chan []byte, publish chan<- []byte) {
	for msg := range messages {
		job, err := parseMessage(msg)
		if err != nil {
			log.Printf("Worker %d: failed to parse message: %v", id, err)
			continue
		}

		processedMessage, err := job.DoJob(ctx, db)
		if err != nil {
			log.Printf("Worker %d: failed to do job: %v", id, err)
			continue
		}

		publish <- processedMessage
	}
}

func StartWorkers(ctx context.Context, db DB, messages <-chan []byte, publish chan<- []byte) {
	for i := 0; i < workerCount; i++ {
		go Worker(ctx, db, i+1, messages, publish)
	}
}
