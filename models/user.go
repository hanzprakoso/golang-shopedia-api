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

type User struct {
	UserID    	int    `gorm:"column:UserID;primary_key" json:"UserID"`
	Username  	string `gorm:"column:Username" json:"Username"`
	FirstName 	string `gorm:"column:FirstName" json:"FirstName"`
	LastName  	string `gorm:"column:LastName" json:"LastName"`
	ProfilePicture  	string `gorm:"column:ProfilePicture" json:"ProfilePicture"`
	Gender    	int    `gorm:"column:Gender" json:"Gender"`
	PhoneNumber	string	`gorm:"column:PhoneNumber" json:"PhoneNumber"`
	Address		string	`gorm:"column:Address" json:"Address"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
