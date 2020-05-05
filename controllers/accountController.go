package controllers

import (
	"fmt"
	"net/http"
	models "shopedia/models"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

type data struct {
	Token     string `json:"token"`
	Message   string `json:"message"`
	Error     error  `json:"error"`
	IsSuccess bool   `json:"is_success"`
}

func (idb *InDB) Login(e echo.Context) error {

	var account models.Account

	db := idb.DB
	email := e.FormValue("email")
	password := e.FormValue("password")

	result := db.Where(models.Account{Email: email, Password: password}).First(&account)

	if result.RecordNotFound() {
		data := data{
			Token:     "",
			Error:     result.Error,
			IsSuccess: !result.RecordNotFound(),
		}
		return e.JSON(http.StatusOK, data)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = account.Email
	claims["admin"] = account.Level > 90
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("aisahaR"))
	if err != nil {
		return err
	}

	data := data{
		Token:     t,
		Error:     result.Error,
		IsSuccess: !result.RecordNotFound(),
	}

	return e.JSON(http.StatusOK, data)
}

func (idb *InDB) RegisterAccount(e echo.Context) error {
	db := idb.DB
	email := e.FormValue("email")
	password := e.FormValue("password")
	account := models.Account{
		Email: email,
		Password: password,
		Level:    1,
	}
	var count int16
	db.Model(&models.Account{}).Where(models.Account{Email: email}).Count(&count)
	fmt.Println(count)
	if count == 0 {
		db.Create(&account)
	} else {
		return e.JSON(http.StatusOK, data{
			Message: "Email already exist",
		})
	}
	return e.JSON(http.StatusOK, account)

}

func (idb *InDB) Coba(e echo.Context) error {
	db := idb.DB
	email := e.FormValue("email")
	password := e.FormValue("password")
	account := models.Account{
		Email: email,
		Password: password,
		Level:    1,
	}
	if db.Select("email").Where(&models.Account{Email: email}).First(&account).RecordNotFound() {
		db.Create(&account)
	} else {
		return e.JSON(http.StatusOK, data{
			Message: "Email already excist",
		})
	}
	return e.JSON(http.StatusOK, account)
}
