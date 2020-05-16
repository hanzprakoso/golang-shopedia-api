package services

import (
	"log"

	"github.com/semicolon27/shopedia-api/database"
	"github.com/semicolon27/shopedia-api/models"
)

func GetVillagesWithSubdistrictBySearch(key string) []models.Village {
	db := database.GetDB()
	var villages []models.Village
	err := db.Where("Name LIKE ?", "%"+key+"%").Find(&villages).Error
	if err != nil {
		log.Fatal(err)
	}
	return villages
}