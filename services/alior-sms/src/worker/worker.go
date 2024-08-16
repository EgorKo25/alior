package worker

import (
	"alior-sms/src/types"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	workerCount int = 4
)

type DB interface {
	InsertService(ctx context.Context, service *types.Service) (int32, error)
	GetPaginatedServices(ctx context.Context, limit, offset int32) ([]*types.Service, error)
}

type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

func doInsertionJob(ctx context.Context, db DB, msg *amqp.Delivery) (*amqp.Publishing, error) {
	var service types.Service

	if msg.Body == nil {
		return nil, fmt.Errorf("message body is nil, msg")
	}

	err := json.Unmarshal(msg.Body, &service)

	if err != nil {
		return nil, err
	}

	// InsertService сама меняет id переданной ей структуре
	_, err = db.InsertService(ctx, &service)
	if err != nil {
		return nil, err
	}

	responseBody, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	response := amqp.Publishing{
		Headers:         msg.Headers,
		ContentType:     msg.ContentType,     // MIME content type
		ContentEncoding: msg.ContentEncoding, // MIME content encoding
		DeliveryMode:    msg.DeliveryMode,    // Transient (0 or 1) or Persistent (2)
		Priority:        msg.Priority,        // 0 to 9
		CorrelationId:   msg.CorrelationId,   // correlation identifier
		ReplyTo:         "ASK",               // address to to reply to (ex: RPC)
		Expiration:      msg.Expiration,      // message expiration spec
		MessageId:       msg.MessageId,       // message identifier
		Timestamp:       time.Now(),          // message timestamp
		Type:            msg.Type,            // message type name
		UserId:          msg.UserId,          // creating user id - ex: "guest"
		AppId:           msg.AppId,           // creating application id
		Body:            responseBody,
	}

	return &response, nil
}

func doPagineJob(ctx context.Context, db DB, msg *amqp.Delivery) (*amqp.Publishing, error) {
	var pagineBody types.PagineBody
	err := json.Unmarshal(msg.Body, &pagineBody)

	if err != nil {
		return nil, err
	}

	// InsertService сама меняет id переданной ей структуре
	services, err := db.GetPaginatedServices(ctx, pagineBody.Limit, pagineBody.Offset)
	if err != nil {
		return nil, err
	}

	responseBody, err := json.Marshal(services)
	if err != nil {
		return nil, err
	}

	response := amqp.Publishing{
		Headers:         msg.Headers,
		ContentType:     msg.ContentType,     // MIME content type
		ContentEncoding: msg.ContentEncoding, // MIME content encoding
		DeliveryMode:    msg.DeliveryMode,    // Transient (0 or 1) or Persistent (2)
		Priority:        msg.Priority,        // 0 to 9
		CorrelationId:   msg.CorrelationId,   // correlation identifier
		ReplyTo:         "ASK",               // address to to reply to (ex: RPC)
		Expiration:      msg.Expiration,      // message expiration spec
		MessageId:       msg.MessageId,       // message identifier
		Timestamp:       time.Now(),          // message timestamp
		Type:            msg.Type,            // message type name
		UserId:          msg.UserId,          // creating user id - ex: "guest"
		AppId:           msg.AppId,           // creating application id
		Body:            responseBody,
	}

	return &response, nil
}

func Worker(ctx context.Context, db DB, id int, writing chan amqp.Publishing, reading chan amqp.Delivery, logger ILogger) {
	var (
		response *amqp.Publishing
		err      error
	)

	for msg := range reading {
		switch msg.Type {
		case "create":
			response, err = doInsertionJob(ctx, db, &msg)
		case "next":
			response, err = doPagineJob(ctx, db, &msg)
		}

		if err != nil {
			logger.Error(fmt.Sprintf("Worker %d: failed to do job: %v", id, err))
			continue
		}

		writing <- *response
	}
}

func StartWorkers(ctx context.Context, db DB, writing chan amqp.Publishing, reading chan amqp.Delivery, logger ILogger) {
	for i := 0; i < workerCount; i++ {
		go Worker(ctx, db, i+1, writing, reading, logger)
	}
}
