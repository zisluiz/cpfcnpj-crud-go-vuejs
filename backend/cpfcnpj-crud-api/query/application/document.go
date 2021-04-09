package application

import (
	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
	"zisluiz.com/cpfcnpj-crud-api/query/dao"
	"zisluiz.com/cpfcnpj-crud-api/query/model"
)

type DocumentQueryApplication struct {
	DocumentDAO dao.DocumentDAO
}

func (app *DocumentQueryApplication) FindDocumentsBy(query *model.Query) *data.Response {
	if query.Filters != nil {
		for _, filter := range query.Filters {
			if filter.Name == "identityNumber" {
				//only for test if input identity number is valid
				_, err := factory.NewIdentity(filter.Value.(string))

				if err != nil {
					return data.ResponseError(err)
				}
			}
		}
	}

	var result = app.DocumentDAO.FindDocumentsBy(query)

	return data.ResponseOk(result)
}
