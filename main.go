package main

import (
	"github.com/semicolon27/shopedia-api/routes"
	"github.com/semicolon27/shopedia-api/database"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	database := database.Connect()
	defer database.Close()

	routes.InitRoutes(e)

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.Fatal(e.Start(":7007"))
}
