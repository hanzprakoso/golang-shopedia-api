package controllers

import (

	"shopedia/models"
	// "github.com/labstack/echo"
)
type data_result struct{
	Data products_result `json:"data"`
}
type products_result struct {
	Products     []models.Product `json:"products"`
	Product 	models.Product
	Message   string `json:"message"`
	Error     error  `json:"error"`
	IsSuccess bool   `json:"is_success"`
}

// func (idb *InDB) getProductByShop(e *echo.Context) {
// 	db := idb.DB
	
// }
