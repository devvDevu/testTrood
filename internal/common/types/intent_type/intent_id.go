package intent_type

type IntentId int

func (ct IntentId) Int() int {
	return int(ct)
}
