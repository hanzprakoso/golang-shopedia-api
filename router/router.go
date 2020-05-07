package router

import (
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {

	//set route here
	AccountRouter(e)
	ProductRouter(e)
}
