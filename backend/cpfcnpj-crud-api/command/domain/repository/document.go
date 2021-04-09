package repository

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
)

type DocumentRepository interface {
	Insert(document *model.Document) *exception.Validations
	Update(document *model.Document) *exception.Validations
	Get(uuid string) (*model.Document, *exception.Validations)
	Delete(uuid string) *exception.Validations
	BlockDocuments(uuids []string, block bool) *exception.Validations
}
