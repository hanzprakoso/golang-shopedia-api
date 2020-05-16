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

type Courier struct {
	CourierID int    `gorm:"column:CourierID;primary_key" json:"CourierID"`
	Name      string `gorm:"column:Name" json:"Name"`
	Slug    string `gorm:"column:Slug" json:"Slug"`
}

// TableName sets the insert table name for this struct type
func (c *Courier) TableName() string {
	return "couriers"
}
