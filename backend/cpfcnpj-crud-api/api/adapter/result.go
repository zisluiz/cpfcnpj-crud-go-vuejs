package adapter

import (
	"zisluiz.com/cnpj-crud-api/application"

	"github.com/labstack/echo/v4"
)

func ToHttpResponse(c echo.Context, result *application.Response) error {
	return c.JSON(result.Status, result.Body)
}
