package intent_repository

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/model/intent_model"

	"github.com/sirupsen/logrus"
)

// this may be agregation request to db
const intentRepositoryName = "IntentRepository"

type IntentRepository struct {
	db mockDbI
}

type mockDbI interface {
	Exec(ctx context.Context, intent string, dest *intent_model.IntentModel) error
	GetList(ctx context.Context) ([]string, error)
	Get(ctx context.Context, intent string, dest *intent_model.IntentModel) error
	Delete(ctx context.Context, intent string) error
}

func NewIntentRepository(db mockDbI) *IntentRepository {
	return &IntentRepository{db: db}
}

func (r *IntentRepository) GetList(ctx context.Context) ([]intent_type.IntentWord, error) {
	const action = "IntentRepository GetList"
	const method = "GetList"
	intents, err := r.db.GetList(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"intent_repository_name": intentRepositoryName,
			"method":                 method,
		}).WithError(err).Error(action)
		return nil, err
	}
	intentWords := make([]intent_type.IntentWord, 0)
	for _, intent := range intents {
		intentWords = append(intentWords, intent_type.IntentWord(intent))
	}
	return intentWords, nil
}

func (r *IntentRepository) Get(ctx context.Context, intent intent_type.IntentWord) (*intent_model.IntentModel, error) {
	dest := new(intent_model.IntentModel)
	err := r.db.Get(ctx, intent.String(), dest)
	logrus.Info(dest)
	return dest, err
}

func (r *IntentRepository) Delete(ctx context.Context, intent intent_type.IntentWord) error {
	return r.db.Delete(ctx, intent.String())
}

func (r *IntentRepository) Exec(ctx context.Context, intent intent_type.IntentWord, data *intent_model.IntentModel) error {
	return r.db.Exec(ctx, intent.String(), data)
}
