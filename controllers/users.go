package controllers

import(
	"fmt"
	// "log"
	"strconv"
	"net/http"

	"github.com/semicolon27/shopedia-api/models"
	"github.com/semicolon27/shopedia-api/services"

	// "github.com/alexedwards/argon2id"
	"github.com/labstack/echo"
)

func RegiterUser(c echo.Context) error {
	gender, _ := strconv.Atoi(c.FormValue("gender"))

	if gender == 0 {
		gender = 3
	}

	var userData = models.User{
		Username : c.FormValue("username"),
		FirstName : c.FormValue("firstname"),
		LastName : c.FormValue("lastname"),
		Gender : gender,
		PhoneNumber : c.FormValue("phonenumber"),
		Address : c.FormValue("address"),
	}

	if services.IsUsernameHasAlreadyTaken(c.FormValue("username")){
		return c.NoContent(http.StatusConflict) //409
	}

	err := services.CreateUserWithResult(userData)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError) //500
	}

	return c.NoContent(http.StatusCreated) //201
}