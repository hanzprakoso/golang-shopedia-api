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

type Invoice struct {
	InvoiceID        int       `gorm:"column:InvoiceID;primary_key" json:"InvoiceID"`
	BuyerUserID      int       `gorm:"column:BuyerUserID" json:"BuyerUserID"`
	CourierServiceID int       `gorm:"column:CourierServiceID" json:"CourierServiceID"`
	DateCreated      time.Time `gorm:"column:DateCreated" json:"DateCreated"`
}

// TableName sets the insert table name for this struct type
func (i *Invoice) TableName() string {
	return "invoice"
}
