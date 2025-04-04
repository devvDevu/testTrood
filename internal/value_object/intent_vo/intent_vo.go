package intent_vo

import (
	"testTrood/internal/common/types/intent_type"
	"testTrood/internal/common/types/user_type"
)

type IntentVo struct {
	userMessage user_type.UserMessage

	intentWord intent_type.IntentWord

	answer intent_type.IntentAnswer
}

func NewIntentVo(userMessage user_type.UserMessage, intentWord intent_type.IntentWord, answer intent_type.IntentAnswer) *IntentVo {
	return &IntentVo{
		userMessage: userMessage,
		intentWord:  intentWord,
		answer:      answer,
	}
}

func (i *IntentVo) GetUserMessage() user_type.UserMessage {
	return i.userMessage
}

func (i *IntentVo) GetIntentWord() intent_type.IntentWord {
	return i.intentWord
}

func (i *IntentVo) GetAnswer() intent_type.IntentAnswer {
	return i.answer
}
