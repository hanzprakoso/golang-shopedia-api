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

type Param struct {
	ParamID int    `gorm:"column:ParamID;primary_key" json:"ParamID"`
	Entity  string `gorm:"column:Entity" json:"Entity"`
	Field   string `gorm:"column:Field" json:"Field"`
	Value   string `gorm:"column:Value" json:"Value"`
}

// TableName sets the insert table name for this struct type
func (p *Param) TableName() string {
	return "params"
}
