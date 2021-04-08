package application

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/repository"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
)

type DocumentCrudApplication struct {
	DocumentRepository repository.DocumentRepository
}

type NewDocumentInputModel struct {
	Name           string `json:"name"`
	IdentityNumber string `json:"identityNumber"`
}

type EditDocumentInputModel struct {
	Uuid string `json:"uuid" swaggerignore:"true"`
	NewDocumentInputModel
}

type DocumentViewModel struct {
	EditDocumentInputModel
	IdentityType string `json:"identityType"`
}

func (app *DocumentCrudApplication) Get(uuid string) *data.Response {
	if len(uuid) == 0 {
		return data.ResponseError(exception.NewValidation("Parameter \"id\" is required!"))
	}

	var document, validations = app.DocumentRepository.Get(uuid)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseOk(&DocumentViewModel{IdentityType: document.Identity.Type(),
		EditDocumentInputModel: EditDocumentInputModel{Uuid: document.Uuid, NewDocumentInputModel: NewDocumentInputModel{Name: document.Name, IdentityNumber: document.Identity.Value()}}})
}

func (app *DocumentCrudApplication) Post(inputDocument *NewDocumentInputModel) *data.Response {
	var document, validations = factory.NewDocument("", inputDocument.Name, inputDocument.IdentityNumber)

	if validations != nil {
		return data.ResponseError(validations)
	}

	validations = app.DocumentRepository.Insert(document)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseCreated()
}

func (app *DocumentCrudApplication) Update(inputDocument *EditDocumentInputModel) *data.Response {
	var document, validations = factory.NewDocument(inputDocument.Uuid, inputDocument.Name, inputDocument.IdentityNumber)

	if validations != nil {
		return data.ResponseError(validations)
	}

	validations = app.DocumentRepository.Update(document)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseOkNoContent()
}

func (app *DocumentCrudApplication) Delete(uuid string) *data.Response {
	if len(uuid) == 0 {
		return data.ResponseError(exception.NewValidation("Parameter \"uuid\" is required!"))
	}

	var validations = app.DocumentRepository.Delete(uuid)

	if validations != nil {
		return data.ResponseError(validations)
	}

	return data.ResponseOkNoContent()
}
