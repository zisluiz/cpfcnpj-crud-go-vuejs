package routes

import (
	"github.com/labstack/echo/v4"
	"zisluiz.com/cpfcnpj-crud-api/command/application"
	"zisluiz.com/cpfcnpj-crud-api/http/adapter"
	"zisluiz.com/cpfcnpj-crud-api/infra/repository"
)

var (
	documentApplication = &application.DocumentCrudApplication{DocumentRepository: &repository.DocumentRepositoryImpl{}}
)

func BindRoute(group *echo.Group) {
	group.GET("/documents", getDocument)
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
