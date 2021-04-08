package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "zisluiz.com/cpfcnpj-crud-api/docs"
	"zisluiz.com/cpfcnpj-crud-api/http/routes"
	appMiddleware "zisluiz.com/cpfcnpj-crud-api/infra/middleware"
)

//To generate refresh docs: swag init -g http/routes/document.go

// @title Swagger CPF/CNPJ API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/

func main() {
	serverStats := appMiddleware.NewStats()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(serverStats.Process)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Stats
	statsRoute := &routes.StatsRoute{Stats: serverStats}
	statsRoute.BindStatsRoute(e)

	// Api
	var group = e.Group("/api")
	routes.BindRoute(group)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))
}
