package controllers

import(
	"fmt"
	// "log"
	"net/http"

	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"

	// "github.com/alexedwards/argon2id"
	"github.com/labstack/echo"
)

func RegisterAccount(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if services.IsEmailHasAlreadyTaken(email) {
		return c.NoContent(http.StatusConflict) //409
	}

	err := services.CreateAccount(email, password)
	if err != nil{
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError) //500
	}

	return c.NoContent(http.StatusCreated) //201
}

func Login(c echo.Context) error {
	var accountData models.Account
	var isAccountExist bool

	email := c.FormValue("email")
	password := c.FormValue("password")

	accountData, isAccountExist = services.GetLoginAccount(email, password)

	if !isAccountExist{
		return c.NoContent(http.StatusUnauthorized) //401
	}

	token, _ := services.CreateLoginToken(accountData)
	return c.JSON(http.StatusOK, token)

}

func CheckUsernameIsAlreadyTaken(c echo.Context) error {
	if services.IsUsernameHasAlreadyTaken(c.Param("username")){
		return c.NoContent(http.StatusConflict) //409
	}
	return c.NoContent(http.StatusOK)
}

func CheckEmailIsAlreadyTaken(c echo.Context) error {
	if services.IsEmailHasAlreadyTaken(c.Param("email")){
		return c.NoContent(http.StatusConflict) //409
	}
	return c.NoContent(http.StatusOK)
}	

func Coba(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func File(c echo.Context) error {
	fileName := c.Param("file_name")
	return c.File(`D:\Download\Videos\`+fileName)
}