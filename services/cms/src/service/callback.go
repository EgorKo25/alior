package service

import (
	"callback_service/src/repository"
	"context"
)

type ICallback interface {
	CreateCallback(ctx context.Context, name string, phone string, callbackType string, idea string) error
}

type CallbackService struct {
	repo repository.IRepository
}

func NewCallbackService(repo repository.IRepository) *CallbackService {
	return &CallbackService{repo: repo}
}

func (s *CallbackService) CreateCallback(ctx context.Context, name string, phone string, callbackType string, idea string) error {
	data := repository.Callback{
		Phone: phone,
		Name:  name,
		Type:  callbackType,
		Idea:  idea,
	}

	return s.repo.CreateCallback(ctx, data)
}
