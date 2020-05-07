package router

import (
	controllers "shopedia/controllers"
	"shopedia/database"

	"github.com/labstack/echo"
)

func AccountRouter(e *echo.Echo) {

	db := database.Connect()      //inisialisasi database
	c := &controllers.InDB{DB: db} // memasukan hasil koneksi database ke controllers

	e.POST("/login", c.Login)
	e.POST("/register", c.RegisterAccount)
	e.POST("/coba", c.Coba)
}
