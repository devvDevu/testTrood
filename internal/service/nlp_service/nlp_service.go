package nlp_service

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/common/types/user_type"
)

type NlpService struct {
	repo nlpRepoI
}

type nlpRepoI interface {
	GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentWord, error)
}

func NewNlpService(repo nlpRepoI) *NlpService {
	return &NlpService{repo: repo}
}

func (s *NlpService) GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentWord, error) {
	return s.repo.GetIntent(ctx, userMessage)
}
