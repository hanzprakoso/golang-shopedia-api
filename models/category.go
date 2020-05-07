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
	CategoryID int    `gorm:"column:CategoryID;primary_key" json:"CategoryID"`
	Name       string `gorm:"column:Name" json:"Name"`
	Domain     string `gorm:"column:Domain" json:"Domain"`
	ParentID   int    `gorm:"column:ParentID" json:"ParentID"`
	CategoryDetails []CategoryDetails `gorm:"foreignkey:CategoryID;ASSOCIATION_FOREIGNKEY:ParentID;"`
}

type CategoryDetails struct {
			CategoryID int     	`gorm:"column:CategoryID"`
			Name       string 	`gorm:"column:Name"`
			Domain		string 	`gorm:"column:Domain"`
			ParentID   int  	`gorm:"column:ParentID"`
}

func (c *CategoryDetails) TableName() string {
	return "categories"
}
// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "categories"
}
