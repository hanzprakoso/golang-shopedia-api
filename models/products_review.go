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

type ProductsReview struct {
	ReviewID  int     `gorm:"column:ReviewID;primary_key" json:"ReviewID"`
	ProductID int     `gorm:"column:ProductID" json:"ProductID"`
	UserID    int     `gorm:"column:UserID" json:"UserID"`
	Rating    float64 `gorm:"column:Rating" json:"Rating"`
	Review    string  `gorm:"column:Review" json:"Review"`
}

// TableName sets the insert table name for this struct type
func (p *ProductsReview) TableName() string {
	return "products_reviews"
}
