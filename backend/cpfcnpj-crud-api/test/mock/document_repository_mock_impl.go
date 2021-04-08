package mock

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/repository"
)

type DocumentRepositoryMockImpl struct {
	repository.DocumentRepository
}

func (repository *DocumentRepositoryMockImpl) Get(uuid string) (*model.Document, *exception.Validations) {
	var validationId *exception.Validations

	if len(uuid) == 0 {
		validationId = exception.NewValidation("Invalid id!")
	}

	document, validation := factory.NewDocument(uuid, "Person Name", "12852042550")

	if validationId != nil && validation != nil {
		validation.Merge(validationId)
	}

	return document, validation
}

func (repository *DocumentRepositoryMockImpl) Delete(uuid string) *exception.Validations {
	var validationId *exception.Validations

	if len(uuid) == 0 {
		validationId = exception.NewValidation("Invalid id!")
	}

	return validationId
}

func (repository *DocumentRepositoryMockImpl) Insert(document *model.Document) *exception.Validations {
	return nil
}

func (repository *DocumentRepositoryMockImpl) Update(document *model.Document) *exception.Validations {
	var validationId *exception.Validations

	if len(document.Uuid) == 0 {
		validationId = exception.NewValidation("Invalid id!")
	}

	return validationId
}
