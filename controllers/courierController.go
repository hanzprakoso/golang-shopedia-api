package controllers

import (
	"fmt"
	"net/http"
	"shopedia/models"
	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"
)

func (idb *InDB) GetCourierWithServices(e echo.Context) error {
	db := idb.DB
	db.LogMode(true)
	// var product models.Product
	// var category models.Category
	// var res products_result
	// db.Model(&category).Related(&product)
	var (
		C models.Courier
		CS models.Courierservice
	)
	// var res result
	// db.Raw("SELECT * FROM products INNER JOIN `categories` ON products.CategoryID = categories.CategoryID").Scan(&res.Product).Scan(&res.Category)
	// db.First(&C, 1)
	if err := db.Model(&C).Preload("Services").Find(&C).Error; err != nil {
		fmt.Println(err)
	  }
	  fmt.Printf("%+v\n", &C)
	db.Model(&C).Related(&CS)
	return e.JSON(http.StatusOK, C)
} 