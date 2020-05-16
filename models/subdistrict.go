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
	Slug        sql.NullString `gorm:"column:Slug" json:"Slug"`
}

type SubdistrictWithCity struct {
	SubDistrictID int            `gorm:"column:SubDistrictID;primary_key" json:"-"`
	CityID        int            `gorm:"column:CityID" json:"-"`
	Name          string         `gorm:"column:Name" json:"Name"`
	Slug        sql.NullString `gorm:"column:Slug" json:"Slug"`
	City	City	`gorm:"foreignkey:CityID;association_foreignkey:CityID"`	
}

// TableName sets the insert table name for this struct type
func (s *Subdistrict) TableName() string {
	return "subdistricts"
}
