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

type Product struct {
	ProductID   int    `gorm:"column:ProductID;primary_key" json:"ProductID"`
	Name        string `gorm:"column:Name" json:"Name"`
	Domain      string `gorm:"column:Domain" json:"Domain"`
	CategoryID  int    `gorm:"column:CategoryID" json:"CategoryID"`
	Description string `gorm:"column:Description" json:"Description"`
	MerchantID  int    `gorm:"column:MerchantID" json:"MerchantID"`
	Price       int    `gorm:"column:Price" json:"Price"`
	Weight      int    `gorm:"column:Weight" json:"Weight"`
	WeightUnit  int    `gorm:"column:WeightUnit" json:"WeightUnit"`
}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}
