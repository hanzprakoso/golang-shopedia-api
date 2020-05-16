package services

import (
	"log"

	"github.com/semicolon27/shopedia-api/database"
	"github.com/semicolon27/shopedia-api/models"
)

func GetSubdistrictsWithCityBySearch(key string) []models.SubdistrictWithCity {
	db := database.GetDB()
	var subdistricts []models.SubdistrictWithCity

	err := db.Table("subdistricts").Where("Name LIKE ?", "%"+key+"%").Preload("City").Find(&subdistricts).Error
	if err != nil {
		log.Fatal(err)
	}

	return subdistricts
}
