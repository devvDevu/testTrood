package intent_model

import "testTrood/internal/common/types/intent_type"

type IntentModel struct {
	Id intent_type.IntentId

	IntentWord intent_type.IntentWord

	Answer intent_type.IntentAnswer
}
