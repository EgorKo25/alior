package service

import (
	"context"
	"time"

	"callback_service/internal/repository"
)

type ICallback interface {
	CreateCallback(ctx context.Context, number string, date string, name string) error
}

type CallbackService struct {
	repo repository.Repository
}

func NewCallbackService(repo repository.Repository) *CallbackService {
	return &CallbackService{repo: repo}
}

func (s *CallbackService) CreateCallback(ctx context.Context, number string, date string, name string) error {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}

	data := repository.Callback{
		Number: number,
		Date:   parsedDate,
		Name:   name,
	}

	return s.repo.CreateCallback(ctx, data)
}
