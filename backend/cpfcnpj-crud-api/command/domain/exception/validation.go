package exception

type rawMessage struct {
	ValidationType string `json:"validationType"`
	Description    string `json:"description"`
}

type rawMessages struct {
	Messages []*rawMessage `json:"messages"`
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

func (a *Validations) ToArrayRaw() *rawMessages {
	rawMessages := rawMessages{Messages: []*rawMessage{}}

	for _, message := range a.Messages {
		rawMessages.Messages = append(rawMessages.Messages, &rawMessage{ValidationType: message.ValidationType(), Description: message.Description()})
	}

	return &rawMessages
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
