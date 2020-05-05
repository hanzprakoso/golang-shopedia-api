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

type Province struct {
	ProvinceID int            `gorm:"column:ProvinceID;primary_key" json:"ProvinceID"`
	Name       string         `gorm:"column:Name" json:"Name"`
	Domain     sql.NullString `gorm:"column:Domain" json:"Domain"`
}

// TableName sets the insert table name for this struct type
func (p *Province) TableName() string {
	return "provinces"
}
