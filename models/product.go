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
	ProductID           int           `gorm:"column:ProductID;primary_key" json:"ProductID"`
	Name                string        `gorm:"column:Name" json:"Name"`
	Slug                string        `gorm:"column:Slug" json:"Slug"`
	CategoryID          int           `gorm:"column:CategoryID" json:"CategoryID"`
	Description         string        `gorm:"column:Description" json:"Description"`
	MerchantID          int           `gorm:"column:MerchantID" json:"MerchantID"`
	DiscountPercentage  int			  `gorm:"column:DiscountPercentage" json:"DiscountPercentage"`
	PriceBeforeDiscount int			  `gorm:"column:PriceBeforeDiscount" json:"PriceBeforeDiscount"`
	Price               int           `gorm:"column:Price" json:"Price"`
	Weight              int           `gorm:"column:Weight" json:"Weight"`
	WeightUnitID        int           `gorm:"column:WeightUnit" json:"-"`
}

type ProductCatalog struct {
	Product
	Rating              float64		`gorm:"column:Rating" json:"Rating"`
	CountReview         int    		`gorm:"column:CountReview" json:"CountReview"`
	Category			[]Category	`json:"Category;"`
	Merchant Merchant 	`gorm:"foreignkey:MerchantID;association_foreignkey:MerchantID"`	
}

type ProductDetail struct {
	Product	struct {
		Product
		WeightUnit string	`gorm:"foreignkey:ParamID;association_foreignkey:WeightUnitID"`
	}
	Category struct {
		Category
		Details []Category
	}
	Merchant Merchant
	ProductsReview []struct{
		ProductsReview
		User	User	`gorm:"foreignkey:UserID;association_foreignkey:UserID"`
	}
	ProductsRating int
}

// TableName sets the insert table name for this struct type
func (p *ProductCatalog) TableName() string {
// func TableName() string {
	return "products"
}
