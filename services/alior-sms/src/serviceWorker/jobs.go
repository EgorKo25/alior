package serviceWorker

import (
	"alior-sms/src/types"
	"context"
	"encoding/json"
)

type PagineJob struct {
	limit  int32
	offset int32
	msg    types.AMQPMessage
}

func (job PagineJob) DoJob(ctx context.Context, db DB) ([]byte, error) {
	services, err := db.GetPaginatedServices(ctx, job.limit, job.offset)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(services)
	job.msg.Body = string(jsonData)

	return jsonData, nil
}

type InsertJob struct {
	service *types.Service
}

func InsertJobMessage(id int32) []byte {
	return []byte("sadsad")
}

func (job InsertJob) DoJob(ctx context.Context, db DB) ([]byte, error) {
	id, err := db.InsertService(ctx, job.service)
	if err != nil {
		return nil, err
	}

	return InsertJobMessage(id), nil
}
