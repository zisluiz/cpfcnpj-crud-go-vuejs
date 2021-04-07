package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"zisluiz.com/cnpj-crud-api/api/routes"
	_ "zisluiz.com/cnpj-crud-api/docs"
)

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
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	var group = e.Group("/api")

	routes.BindRoute(group)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))
}
