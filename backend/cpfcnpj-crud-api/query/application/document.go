package application

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
)

func GetDocument(uuid string) *data.Response {
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
