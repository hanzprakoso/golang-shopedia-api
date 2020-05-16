package routes

import (
	c "github.com/semicolon27/shopedia-api/controllers"
	"github.com/labstack/echo"
)

func ProductRouter(e *echo.Echo){
	e.GET("/search/products", c.SearchProducts)
	e.GET("/:merchant_slug/:product_slug", c.DetailsProductByMerchantProductSlug)
	// e.GET("/gg", c.DetailsProduct)
}