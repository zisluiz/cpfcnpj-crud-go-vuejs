package routes

import (
	"github.com/labstack/echo/v4"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/http/adapter"
	"zisluiz.com/cpfcnpj-crud-api/infra/dao"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
	"zisluiz.com/cpfcnpj-crud-api/query/application"
	"zisluiz.com/cpfcnpj-crud-api/query/model"
)

var (
	queryDocumentApplication = &application.DocumentQueryApplication{DocumentDAO: &dao.DocumentDAOImpl{}}
)

func BindQueryRoute(group *echo.Group) {
	group.GET("/documents", queryDocuments)
}

// GetDocument godoc
// @Summary Get an document
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} application.DocumentViewModel
// @Failure 400 {object} data.Response
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/documents [get]
func queryDocuments(c echo.Context) error {
	var query = new(model.Query)

	if err := c.Bind(query); err != nil {
		return adapter.ToHttpResponse(c, data.ResponseError(exception.NewError(err.Error())))
	}

	query.ParseJsonParams()

	var response = queryDocumentApplication.FindDocumentsBy(query)
	return adapter.ToHttpResponse(c, response)
}
