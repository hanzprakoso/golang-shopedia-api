package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"shopedia/models"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
)

func product(db *gorm.DB, w string, p ...string) models.Product {
	var (
		P models.Product
		C models.Category
		CD models.CategoryDetails 
	)
	where := fmt.Sprintf("%v", w)
	param := p
	db.Preload("Category").Preload("Merchant").Preload("WeightUnit").Find(&P, where, param )

	id := P.Category.ParentID
	db.Where("CategoryID = ?", id).Find(&CD)
	P.Category.CategoryDetails = append(P.Category.CategoryDetails, CD)
	var i = 0
	for {
		CD = models.CategoryDetails{}
		if id := P.Category.CategoryDetails[i].ParentID; id != 0{
			db.Model(&models.Category{}).Raw("SELECT * FROM categories WHERE CategoryID = ?", id).Scan(&CD)
			P.Category.CategoryDetails = append(P.Category.CategoryDetails, CD)
		}else{
			break
		}
		i++
	}
	fmt.Printf("%+v\n", &C)
	return P
}
func (idb *InDB) GetProducts(e echo.Context) error {
	db := idb.DB
	db.LogMode(true)
	//fungsi produk dengan parameter kedua WHERE dan ketiga parameternya
	//product(db, "1", "1") => ...WHERE "1" = "1"q v
	return e.JSON(http.StatusOK, product(db, "1 = ?", "1"))
} 

func (idb *InDB) GetProductsByMerchantID(e echo.Context) error {
	db := idb.DB
	db.LogMode(true)
	MID := e.Param("merchant_id")
	//fungsi produk dengan parameter kedua WHERE dan ketiga parameternya
	//product(db, "1", "1") => ...WHERE "1" = "1"q v
	return e.JSON(http.StatusOK, product(db, "MerchantID = ?", MID))
} 

func (idb *InDB) GetProductsByMerchantName(e echo.Context) error {
	db := idb.DB
	MN := e.Param("merchant_name")
	var m models.Merchant
	
	db.LogMode(true)
	MN = strings.Replace(MN, "-", " ", -1)  //replace all stripes with spaces

	//fungsi produk dengan parameter kedua WHERE dan ketiga parameternya
	//product(db, "1", "1") => ...WHERE "1" = "1"q v
	db.Select("MerchantID").Find(&m, "Name = ?", MN)
	return e.JSON(http.StatusOK, product(db, "MerchantID = ?", strconv.Itoa(m.MerchantID)))
} 

func (idb *InDB) GetProductByBothNames(e echo.Context) error {
	db := idb.DB
	MN := e.Param("merchant_name")
	PD := e.Param("product_domain")
	var m models.Merchant
	
	db.LogMode(true)
	MN = strings.Replace(MN, "-", " ", -1)  //replace all stripes with spaces

	//fungsi produk dengan parameter kedua WHERE dan ketiga parameternya
	//product(db, "1", "1") => ...WHERE "1" = "1"q v
	db.Select("MerchantID").Find(&m, "Name = ?", MN)
	return e.JSON(http.StatusOK, product(db, fmt.Sprintf("MerchantID = %v AND Domain = ?", strconv.Itoa(m.MerchantID)), PD))
}

func (idb *InDB) Category(e echo.Context) error {
	db := idb.DB
	var c models.Category
	db.LogMode(true)

	db.Preload("Parents").Find(&c)
	return e.JSON(http.StatusOK, c)
}