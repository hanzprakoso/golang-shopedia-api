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

type Invoiceitem struct {
	InvoiceItemID int `gorm:"column:InvoiceItemID;primary_key" json:"InvoiceItemID"`
	InvoiceID     int `gorm:"column:InvoiceID" json:"InvoiceID"`
	ProductID     int `gorm:"column:ProductID" json:"ProductID"`
	Quantity      int `gorm:"column:Quantity" json:"Quantity"`
}

// TableName sets the insert table name for this struct type
func (i *Invoiceitem) TableName() string {
	return "invoiceitem"
}
