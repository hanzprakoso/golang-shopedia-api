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

type Courierservice struct {
	CourierServiceID int    `gorm:"column:CourierServiceID;primary_key" json:"CourierServiceID"`
	CourierID        int    `gorm:"column:CourierID" json:"CourierID"`
	Name             string `gorm:"column:Name" json:"Name"`
	Domain           string `gorm:"column:Domain" json:"Domain"`
}

// TableName sets the insert table name for this struct type
func (c *Courierservice) TableName() string {
	return "courierservices"
}
