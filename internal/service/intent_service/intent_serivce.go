package intent_service

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/model/intent_model"
)

type IntentService struct {
	repo intentRepoI
}

type intentRepoI interface {
	GetList(ctx context.Context) ([]intent_type.IntentWord, error)
	Get(ctx context.Context, intent intent_type.IntentWord) (*intent_model.IntentModel, error)
	Delete(ctx context.Context, intent intent_type.IntentWord) error
	Exec(ctx context.Context, intent intent_type.IntentWord, data *intent_model.IntentModel) error
}

func NewIntentService(repo intentRepoI) *IntentService {
	return &IntentService{repo: repo}
}

func (s *IntentService) GetList(ctx context.Context) ([]intent_type.IntentWord, error) {
	return s.repo.GetList(ctx)
}

func (s *IntentService) Get(ctx context.Context, intent intent_type.IntentWord) (*intent_model.IntentModel, error) {
	return s.repo.Get(ctx, intent)
}

func (s *IntentService) Delete(ctx context.Context, intent intent_type.IntentWord) error {
	return s.repo.Delete(ctx, intent)
}

func (s *IntentService) Exec(ctx context.Context, intent intent_type.IntentWord, data *intent_model.IntentModel) error {
	return s.repo.Exec(ctx, intent, data)
}
