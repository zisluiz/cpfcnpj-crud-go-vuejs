package application

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/repository"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
)

type BlockListApplication struct {
	DocumentRepository repository.DocumentRepository
}

func (app *BlockListApplication) BlockDocuments(uuids []string) *data.Response {
	var validations *exception.Validations

	if uuids == nil {
		return data.ResponseError(exception.NewValidation("At least one uuid is required!"))
	}

	for _, uuid := range uuids {
		if len(uuid) == 0 {
			if validations == nil {
				validations = exception.NewValidations()
			}

			validations.AddValidation("Empty uuid is not allowed!")
		}
	}

	if validations != nil {
		return data.ResponseError(validations)
	}

	validations = app.DocumentRepository.BlockDocuments(uuids, true)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseOkNoContent()
}

func (app *BlockListApplication) UnBlockDocuments(uuids []string) *data.Response {
	var validations *exception.Validations

	if uuids == nil {
		return data.ResponseError(exception.NewValidation("At least one uuid is required!"))
	}

	for _, uuid := range uuids {
		if len(uuid) == 0 {
			validations.AddValidation("Empty uuid is not allowed!")
		}
	}

	if validations != nil {
		return data.ResponseError(validations)
	}

	validations = app.DocumentRepository.BlockDocuments(uuids, false)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseOkNoContent()
}
