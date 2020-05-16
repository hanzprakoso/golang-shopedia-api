package services

import (
	"fmt"
	"log"
	"time"

	"github.com/semicolon27/shopedia-api/database"
	"github.com/semicolon27/shopedia-api/models"

	"github.com/dgrijalva/jwt-go"
	argon "github.com/alexedwards/argon2id"
)

func IsEmailHasAlreadyTaken(email string) bool {
	db := database.GetDB()
	var account models.Account

	return !db.Select("Email").Where(&models.Account{Email: email}).First(&account).RecordNotFound()
}

func IsUsernameHasAlreadyTaken(username string) bool {
	db := database.GetDB()
	var user models.User

	return !db.Select("Username").Where(&models.User{Username: username}).First(&user).RecordNotFound()
}

func CreateAccount(email string, password string) error {
	db := database.GetDB()
	var account = models.Account{Email: email, Password: encryptPassword(password), UserID: 404, Level: 4}
	
	return db.Create(&account).Error
}

func GetLoginAccount(email string, password string) (models.Account, bool) {
	var account models.Account
	db := database.GetDB()

	db.Where(models.Account{Email: email}).First(&account)
	isPasswordMatch, err := argon.ComparePasswordAndHash(password, account.Password)

	if !isPasswordMatch || err != nil{
		fmt.Println(err)
		return account, false
	}

	return account, true
}

func CreateLoginToken(account models.Account) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = account.Email
	claims["admin"] = account.Level > 90
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as return.
	return token.SignedString([]byte("aisahaR"))
}

func encryptPassword(password string) string {
	hash, err := argon.CreateHash(password, argon.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}