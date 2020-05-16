package routes

import(
	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo){

	ProductRouter(e)
	AccountRouter(e)

}