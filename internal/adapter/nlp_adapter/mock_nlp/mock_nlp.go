package mock_nlp

import (
	"context"
	"math/rand"
	"time"
)

// this may be adapter for external nlp service
type MockNLP struct {
}

func NewMockNLP() *MockNLP {
	return &MockNLP{}
}

func (m *MockNLP) GetIntent(ctx context.Context, text string) (string, error) {
	intents := []string{"change_password", "change_email"}
	time.Sleep(time.Second * 2)
	return intents[rand.Intn(len(intents))], nil
}
