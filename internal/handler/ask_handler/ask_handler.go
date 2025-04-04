package ask_handler

import (
	"context"
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/common/types/user_type"
	"testTrood/internal/value_object/intent_vo"

	"github.com/sirupsen/logrus"
)

const handlerName = "ask_handler"

type AskHandler struct {
	usecase askUsecaseI
}

type askUsecaseI interface {
	GetIntent(ctx context.Context, userMessage user_type.UserMessage) (intent_type.IntentId, *intent_vo.IntentVo, error)
}

func NewAskHandler(usecase askUsecaseI) *AskHandler {
	return &AskHandler{usecase: usecase}
}

func (h *AskHandler) Handle(ctx context.Context, userMessage user_type.UserMessage) {
	const action = "AskHandler Handle"
	const method = "Handle"

	logrus.Info(action)
	id, intentVo, err := h.usecase.GetIntent(ctx, userMessage)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler_name": handlerName,
			"method":       method,
		}).WithError(err).Error(action)
		return
	}
	logrus.Infof("user_message: %s, intent_id: %d, intent_vo: %v", userMessage, id, intentVo)
}
