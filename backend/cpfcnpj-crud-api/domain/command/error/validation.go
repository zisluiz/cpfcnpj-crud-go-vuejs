package error

type Validation struct {
	message string
}

func NewValidation(message string) *Validation {
	return &Validation{message}
}

func (a *Validation) Message() string {
	return a.message
}
