package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"zisluiz.com/cnpj-crud-api/api/adapter"
	"zisluiz.com/cnpj-crud-api/application"
)

func BindRoute(group *echo.Group) {
	group.POST("/documents", postDocument)
	group.GET("/documents", getDocument)
}

// PostDocument godoc
// @Summary Create an document
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param name body application.NewDocumentInput true "Add name"
// @Success 200
// @Failure 400 {object} application.Response
// @Failure 404 {object} application.Response
// @Failure 500 {object} application.Response
// @Router /api/documents [post]

//To generate refresh docs: swag init -g api/routes/document.go
func postDocument(c echo.Context) error {
	var input = new(application.NewDocumentInput)

	if err := c.Bind(input); err != nil {
		return err
	}

	var response = application.PostDocument(input)

	return adapter.ToHttpResponse(c, response)
}

func getDocument(c echo.Context) error {

	return c.String(http.StatusOK, "Hello")
}
