package user_type

type UserMessage string

func (ct UserMessage) String() string {
	return string(ct)
}
