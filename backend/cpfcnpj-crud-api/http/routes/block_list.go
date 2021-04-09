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
	blockListApplication = &application.BlockListApplication{DocumentRepository: &repository.DocumentRepositoryImpl{}}
)

func BindRouteBlock(group *echo.Group) {
	group.PUT("/blocklist", putBlockList)
	group.PUT("/unblocklist", putUnBlockList)
}

type BlockListInputModel struct {
	Uuids []string `json:"uuids"`
}

// Blocklist godoc
// @Summary Block a list of documents
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param ids body string true "Document ids"
// @Success 204
// @Failure 400 {object} data.Response
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/blocklist [put]
func putBlockList(c echo.Context) error {
	var input = new(BlockListInputModel)

	if err := c.Bind(input); err != nil {
		return adapter.ToHttpResponse(c, data.ResponseError(exception.NewError(err.Error())))
	}

	var response = blockListApplication.BlockDocuments(input.Uuids)
	return adapter.ToHttpResponse(c, response)
}

// Unblocklist godoc
// @Summary Unblocklist  a list of documents
// @Description add by json document
// @Tags documents
// @Accept  json
// @Produce  json
// @Param ids body string true "Document ids"
// @Success 204
// @Failure 400 {object} data.Response
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /api/unblocklist [put]
func putUnBlockList(c echo.Context) error {
	var input = new(BlockListInputModel)

	if err := c.Bind(input); err != nil {
		return adapter.ToHttpResponse(c, data.ResponseError(exception.NewError(err.Error())))
	}

	var response = blockListApplication.UnBlockDocuments(input.Uuids)
	return adapter.ToHttpResponse(c, response)
}
