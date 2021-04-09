package dao

import (
	"zisluiz.com/cpfcnpj-crud-api/query/model"
)

type DocumentDAO interface {
	FindDocumentsBy(query *model.Query) *model.Result
}
