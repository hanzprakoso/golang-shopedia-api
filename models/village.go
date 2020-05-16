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

type Village struct {
	VillageID     int    `gorm:"column:VillageID;primary_key" json:"VillageID"`
	SubDistrictID int    `gorm:"column:SubDistrictID" json:"SubDistrictID"`
	Name          string `gorm:"column:Name" json:"Name"`
	Slug        string `gorm:"column:Slug" json:"Slug"`
	PostalCode    string `gorm:"column:PostalCode" json:"PostalCode"`
}

// TableName sets the insert table name for this struct type
func (v *Village) TableName() string {
	return "villages"
}
