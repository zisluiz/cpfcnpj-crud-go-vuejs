package routes

import (
	"github.com/labstack/echo/v4"
	"zisluiz.com/cpfcnpj-crud-api/command/application"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/http/adapter"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
	"zisluiz.com/cpfcnpj-crud-api/infra/repository"
)

var (
	documentApplication = &application.DocumentCrudApplication{DocumentRepository: &repository.DocumentRepositoryImpl{}}
)

func BindRoute(group *echo.Group) {
	group.POST("/documents", postDocument)
	group.GET("/documents/:id", getDocument)
	group.PUT("/documents/:id", putDocument)
	group.DELETE("/documents/:id", deleteDocument)
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
// @Router /api/documents/{id} [get]
func getDocument(c echo.Context) error {
	id := c.Param("id")

	var response = documentApplication.Get(id)
	return adapter.ToHttpResponse(c, response)
}

// PostDocument godoc
// @Summary Create an document
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param Document body application.NewDocumentInputModel true "Add new document"
// @Success 201
// @Failure 400 {object} data.Response
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/documents [post]
func postDocument(c echo.Context) error {
	var input = new(application.NewDocumentInputModel)

	if err := c.Bind(input); err != nil {
		return adapter.ToHttpResponse(c, data.ResponseError(exception.NewError(err.Error())))
	}

	var response = documentApplication.Post(input)
	return adapter.ToHttpResponse(c, response)
}

// PutDocument godoc
// @Summary Create an document
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param Document body application.EditDocumentInputModel true "Update document"
// @Success 204
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/documents/{id} [put]
func putDocument(c echo.Context) error {
	var input = new(application.EditDocumentInputModel)

	if err := c.Bind(input); err != nil {
		return adapter.ToHttpResponse(c, data.ResponseError(exception.NewError(err.Error())))
	}

	var response = documentApplication.Update(input)
	return adapter.ToHttpResponse(c, response)
}

// DeleteDocument godoc
// @Summary Delete an document
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 204
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/documents/{id} [delete]
func deleteDocument(c echo.Context) error {
	id := c.Param("id")

	var response = documentApplication.Delete(id)
	return adapter.ToHttpResponse(c, response)
}
