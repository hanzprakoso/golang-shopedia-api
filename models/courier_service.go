package models

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type CourierService struct {
	CourierServiceID int    `gorm:"column:CourierServiceID;primary_key" json:"CourierServiceID"`
	CourierID        int    `gorm:"column:CourierID" json:"CourierID"`
	Name             string `gorm:"column:Name" json:"Name"`
	Slug           string `gorm:"column:Slug" json:"Slug"`
}

// TableName sets the insert table name for this struct type
func (c *CourierService) TableName() string {
	return "courier_services"
}
