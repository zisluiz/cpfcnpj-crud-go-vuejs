package adapter

import (
	"encoding/json"

	"zisluiz.com/cpfcnpj-crud-api/infra/data"

	"github.com/labstack/echo/v4"
)

func ToHttpResponse(c echo.Context, result *data.Response) error {
	if result.Body != nil {
		var json, _ = json.Marshal(result.Body)
		return c.JSONBlob(result.Status, json)
	} else {
		return c.NoContent(result.Status)
	}
}
