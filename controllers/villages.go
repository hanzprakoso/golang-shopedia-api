package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"
)

func SearchVillages(c echo.Context) error {

	var villages [] models.Village

	key := c.QueryParam("key")
	
	villages = services.GetVillagesWithSubdistrictBySearch(key)
	var res = map[string][]models.Village{
		"data": villages,
	}
	return c.JSON(http.StatusOK, res)
}