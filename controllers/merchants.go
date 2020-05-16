package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"
)

// func AddMerchant(c echo.Context) error {
// 	var merchantData = models.Merchant{
// 		Name
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

func SearchMerchantByName(c echo.Context) error {
	var merchants = services.GetMerchantByName(c.Param("merchant_name"))

	var res = map[string][]models.Merchant{
		"data": merchants,
	}
	
	return c.JSON(http.StatusOK, res)
}