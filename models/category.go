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

type Category struct {
	CategoryID    int           `gorm:"column:CategoryID;primary_key" json:"CategoryID"`
	Name          string        `gorm:"column:Name" json:"Name"`
	Slug          string        `gorm:"column:Slug" json:"Slug"`
	ParentID      sql.NullInt64 `gorm:"column:ParentID" json:"-"`
	GrandParentID sql.NullInt64 `gorm:"column:GrandParentID" json:"-"`
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "categories"
}
