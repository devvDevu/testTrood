package intent_usecase

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/common/types/user_type"
	"testTrood/internal/model/intent_model"
	"testTrood/internal/value_object/intent_vo"

	"github.com/sirupsen/logrus"
)

const usecaseName = "IntentUsecase"

type IntentUsecase struct {
	service    intentServiceI
	nlpService nlpServiceI
}

type intentServiceI interface {
	Get(ctx context.Context, intent intent_type.IntentWord) (*intent_model.IntentModel, error)
}

type nlpServiceI interface {
	GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentWord, error)
}

func NewIntentUsecase(service intentServiceI, nlpService nlpServiceI) *IntentUsecase {
	return &IntentUsecase{service: service, nlpService: nlpService}
}

func (u *IntentUsecase) GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentId, *intent_vo.IntentVo, error) {
	const action = "IntentUsecase GetIntent "
	const method = "GetIntent"

	intent, err := u.nlpService.GetIntent(ctx, userMessage)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"usecase_name": usecaseName,
			"method":       method,
		}).WithError(err).Error(action)
		return 0, nil, err
	}
	intentModel, err := u.service.Get(ctx, intent)
	logrus.Info(intentModel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"usecase_name": usecaseName,
			"method":       method,
		}).WithError(err).Error(action)
		return 0, nil, err
	}
	intentVo := intent_vo.NewIntentVo(userMessage, intentModel.IntentWord, intentModel.Answer)
	return intentModel.Id, intentVo, nil
}
