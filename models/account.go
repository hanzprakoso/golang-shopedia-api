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

type Account struct {
	AccountID int    `gorm:"column:AccountID;primary_key" json:"AccountID"`
	Email     string `gorm:"column:Email" json:"Email"`
	Password  string `gorm:"column:Password" json:"Password"`
	UserID    int    `gorm:"column:UserID" json:"UserID"`
	Level     int    `gorm:"column:Level" json:"Level"`
}

// TableName sets the insert table name for this struct type
func (a *Account) TableName() string {
	return "accounts"
}
