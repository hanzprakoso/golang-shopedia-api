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

type City struct {
	CityID     int            `gorm:"column:CityID;primary_key" json:"CityID"`
	ProvinceID int            `gorm:"column:ProvinceID" json:"ProvinceID"`
	Name       string         `gorm:"column:Name" json:"Name"`
	Domain     sql.NullString `gorm:"column:Domain" json:"Domain"`
}

// TableName sets the insert table name for this struct type
func (c *City) TableName() string {
	return "cities"
}
