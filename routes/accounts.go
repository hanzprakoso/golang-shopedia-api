package routes

import (
	c "github.com/semicolon27/shopedia-api/controllers"
	"github.com/labstack/echo"
)

func AccountRouter(e *echo.Echo){
	e.POST("/login", c.Login)
	e.POST("/register", c.RegisterAccount)
	e.POST("/email/isexist/:email", c.CheckEmailIsAlreadyTaken)
	e.POST("/username/isexist/:username", c.CheckUsernameIsAlreadyTaken)
	e.POST("/coba", c.Coba)
	e.GET("/file/:file_name", c.File)
}