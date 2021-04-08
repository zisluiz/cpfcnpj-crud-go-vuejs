package exception

type rawMessage struct {
	ValidationType string
	Description    string
}

type message interface {
	ValidationType() string
	Description() string
}

type validation struct {
	description string
}

type error struct {
	description string
}

type Validations struct {
	Messages []message
}

func (a *error) ValidationType() string {
	return "error"
}

func (a *validation) ValidationType() string {
	return "validation"
}

func (a *validation) Description() string {
	return a.description
}

func (a *error) Description() string {
	return a.description
}

func (a *Validations) AddValidation(message string) {
	a.Messages = append(a.Messages, &validation{description: message})
}

func (a *Validations) AddError(message string) {
	a.Messages = append(a.Messages, &error{description: message})
}

func (a *Validations) Merge(validations *Validations) {
	a.Messages = append(a.Messages, validations.Messages...)
}

func (a *Validations) ToArrayRaw() []rawMessage {
	raw := []rawMessage{}

	for _, message := range a.Messages {
		raw = append(raw, rawMessage{ValidationType: message.ValidationType(), Description: message.Description()})
	}

	return raw
}

func NewValidations() *Validations {
	return &Validations{}
}

func NewValidation(message string) *Validations {
	var validations = NewValidations()
	validations.AddValidation(message)

	return validations
}

func NewError(message string) *Validations {
	var validations = NewValidations()
	validations.AddError(message)

	return validations
}
