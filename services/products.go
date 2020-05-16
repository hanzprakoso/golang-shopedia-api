package services

import (
	// "fmt"
	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/database"
	// "github.com/gosimple/slug"
)

type productReviewsByProductIDStruct struct {
	ProductID int `gorm:"column:ProductID" json:"ProductID"`
	CountReview int `gorm:"column:CountReview" json:"CountReview"`
	Rating float64 `gorm:"column:Rating" json:"Rating"`
}

func GetProductsCatalogBySearch(key string, page int, rows int, order_by string) []models.ProductCatalog {
	db := database.GetDB()
	var products []models.ProductCatalog
	var productReviews []productReviewsByProductIDStruct

	switch {
	case rows == 0:
		rows = 20
		fallthrough
	case (order_by != "Name" || order_by != "ProductID" || order_by != "CountReview" || order_by != "Rating"):
		order_by = "Name"
	}

	db.Preload("Merchant").Preload("Merchant.City").Limit(rows).Order(order_by).Find(&products, "Name LIKE ?", "%"+key+"%")

	productIDs := getProductIDFromProductCatalog(products)

	db.Raw("SELECT ProductID, COUNT(Review) AS CountReview, AVG(Rating) AS Rating FROM products_reviews WHERE ProductID IN (?)", productIDs).Scan(&productReviews)
	for i , productSlug := range products {
		for j, r := range productReviews {
			if productSlug.ProductID == r.ProductID {
				products[i].CountReview = r.CountReview
				products[i].Rating = r.Rating
				productReviews = productReviews[j+1:] 	//memotong array "productReviews" agar tidak menggunakan kembali productid yang telah digunakan
				break
			}
		}
	}

	return products
}

func getProductIDFromProductCatalog(products []models.ProductCatalog) []int {
	var productID []int
	for _, productSlug := range products {
		productID = append(productID, productSlug.ProductID)
	}
	return productID
}

func GetProductDetails(productSlug string, merchantSlug string) models.ProductDetail {
	db := database.GetDB()
	var details models.ProductDetail

	db.Table("products").Where("Slug = ?", productSlug).Preload("WeightUnit").Find(&details.Product)		//get product
	db.Table("merchants").Where("Slug = ?", merchantSlug).Preload("City").Find(&details.Merchant)		//get merchant
	db.Table("categories").Where("CategoryID = ?", details.Product.CategoryID).Find(&details.Category)		//get category
	db.Raw("SELECT * FROM categories WHERE CategoryID IN ((SELECT ParentID FROM categories WHERE CategoryID = ?),(SELECT GrandParentID FROM categories WHERE CategoryID = ?), ?)", details.Product.CategoryID, details.Product.CategoryID, details.Product.CategoryID).
	   Scan(&details.Category.Details) // get category's details 
	db.Preload("User").Where("ProductID = ?", details.Product.ProductID).Find(&details.ProductsReview) //get review and user

	return details
}