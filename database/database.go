package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/shopedia?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		return nil
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return db
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}