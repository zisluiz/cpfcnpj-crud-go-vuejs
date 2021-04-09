package factory

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
)

func NewDocument(uuid string, name string, identityNumber string) (*model.Document, *exception.Validations) {
	var validations *exception.Validations

	if len(name) < 4 {
		validations = exception.NewValidation("Name must have at least 3 characters!")
	}

	var identity, validationIdentity = NewIdentity(identityNumber)
	var document = model.NewDocument(uuid, name, identity)

	if validationIdentity != nil {
		if validations != nil {
			validations.Merge(validationIdentity)
		} else {
			validations = validationIdentity
		}
	}

	return document, validations
}

func NewDocumentWithBlocked(uuid string, name string, identityNumber string, blocked bool) (*model.Document, *exception.Validations) {
	document, validations := NewDocument(uuid, name, identityNumber)

	if document != nil {
		document.Blocked = blocked
	}

	return document, validations
}
