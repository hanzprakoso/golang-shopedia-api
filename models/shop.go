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

type Shop struct {
	ShopID         int    `gorm:"column:ShopID;primary_key" json:"ShopID"`
	Name           string `gorm:"column:Name" json:"Name"`
	Description    string `gorm:"column:Description" json:"Description"`
	OwnerUserID    int    `gorm:"column:OwnerUserID" json:"OwnerUserID"`
	LocationCityID int    `gorm:"column:LocationCityID" json:"LocationCityID"`
	Address        string `gorm:"column:Address" json:"Address"`
}

// TableName sets the insert table name for this struct type
func (s *Shop) TableName() string {
	return "shops"
}
