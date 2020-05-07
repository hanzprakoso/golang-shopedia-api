package router

import (
	"shopedia/controllers"
	"shopedia/database"

	"github.com/labstack/echo"
)

func ProductRouter(e *echo.Echo) {

	db := database.Connect()      //inisialisasi database
	c := &controllers.InDB{DB: db} // memasukan hasil koneksi database ke controllers

	e.GET("/p", c.GetProducts)
	e.GET("/p/merchant/id/:merchant_id", c.GetProductsByMerchantID)
	e.GET("/p/merchant/:merchant_name", c.GetProductsByMerchantName)
	e.GET("/p/:merchant_name/:product_domain", c.GetProductByBothNames)
	e.GET("/c", c.Category)
	
}
