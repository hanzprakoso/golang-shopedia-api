package services

import (
	"log"

	"github.com/semicolon27/shopedia-api/database"
	"github.com/semicolon27/shopedia-api/models"
)

func GetMerchantByName(merchantName string) []models.Merchant {
	db := database.GetDB()
	var merchants []models.Merchant

	err := db.Where("Name LIKE ?", "%"+merchantName+"%").Find(&merchants).Error
	if err != nil {
		log.Fatal(err)
	}

	return merchants
}
