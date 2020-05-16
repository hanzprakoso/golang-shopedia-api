package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"
)

func SearchProducts(c echo.Context) error {

	var products [] models.ProductCatalog

	key := c.QueryParam("key")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	rows, _ := strconv.Atoi(c.QueryParam("rows"))
	order_by := c.QueryParam("ob")

	products = services.GetProductsCatalogBySearch(key, page, rows, order_by)
	var res = map[string][]models.ProductCatalog{
		"data": products,
	}
	return c.JSON(http.StatusOK, res)
}

func DetailsProductByMerchantProductSlug(c echo.Context) error {
	product_slug := c.Param("product_slug")
	merchant_slug := c.Param("merchant_slug")

	return c.JSON(http.StatusOK, services.GetProductDetails(product_slug, merchant_slug))
}