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

type Subdistrict struct {
	SubDistrictID int            `gorm:"column:SubDistrictID;primary_key" json:"SubDistrictID"`
	CityID        int            `gorm:"column:CityID" json:"CityID"`
	Name          string         `gorm:"column:Name" json:"Name"`
	Domain        sql.NullString `gorm:"column:Domain" json:"Domain"`
}

// TableName sets the insert table name for this struct type
func (s *Subdistrict) TableName() string {
	return "subdistricts"
}
