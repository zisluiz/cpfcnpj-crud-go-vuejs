package routes

import (
	"github.com/labstack/echo/v4"
	"zisluiz.com/cpfcnpj-crud-api/http/adapter"
	"zisluiz.com/cpfcnpj-crud-api/infra/data"
	"zisluiz.com/cpfcnpj-crud-api/infra/middleware"
)

type StatsRoute struct {
	Stats *middleware.Stats
}

func (s *StatsRoute) BindStatsRoute(echo *echo.Echo) {
	echo.GET("/stats", s.getStats)
}

// Stats server godoc
// @Summary Get an document
// @Description add by json document
// @Tags stats
// @Accept  json
// @Produce  json
// @Success 200 {object} data.Response
// @Failure 400 {object} data.Response
// @Failure 404 {object} data.Response
// @Failure 500 {object} data.Response
// @Router /stats [get]
func (s *StatsRoute) getStats(c echo.Context) error {
	s.Stats.Mutex.RLock()
	defer s.Stats.Mutex.RUnlock()

	return adapter.ToHttpResponse(c, data.ResponseOk(s))
}
