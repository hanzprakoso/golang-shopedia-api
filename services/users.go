package services

import (
	"github.com/semicolon27/shopedia-api/database"
	"github.com/semicolon27/shopedia-api/models"
)

func CreateUserWithResult(userData models.User) error {	
	db := database.GetDB()
	return db.Create(&userData).Error
}