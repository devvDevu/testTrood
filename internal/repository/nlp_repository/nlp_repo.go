package nlp_repository

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/common/types/user_type"
)

const nlpRepositoryName = "NlpRepository"

type NlpRepository struct {
	adapter nlpAdapterI
}

type nlpAdapterI interface {
	GetIntent(ctx context.Context, text string) (string, error)
}

func NewNlpRepository(adapter nlpAdapterI) *NlpRepository {
	return &NlpRepository{adapter: adapter}
}

func (r *NlpRepository) GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentWord, error) {
	intent, err := r.adapter.GetIntent(ctx, userMessage.String())
	return intent_type.IntentWord(intent), err
}
