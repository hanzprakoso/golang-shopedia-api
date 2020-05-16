package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"
)

func SearchSubdistricts(c echo.Context) error {
	var subdistricts [] models.SubdistrictWithCity

	key := c.QueryParam("key")
	
	subdistricts = services.GetSubdistrictsWithCityBySearch(key)
	
	var res = map[string][]models.SubdistrictWithCity{
		"data": subdistricts,
	}

	return c.JSON(http.StatusOK, res)
}
