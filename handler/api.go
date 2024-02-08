package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/totoledao/brew-journal/handler/utils"
)

type ReqError struct {
  Message  string `json:"message" xml:"message"`
  Error  string `json:"error" xml:"error"`
}

type UserCredentials struct {
	ID       *string
	Username *string
}

func errHelper(c echo.Context, message string, err error){	
		fmt.Println(message, err)

		e := &ReqError{
			Message:  message,
			Error:  err.Error(),
		}
  	c.JSON(http.StatusInternalServerError, e)
}

func handleNextScreenLogin(c echo.Context, u *UserCredentials) error {

	err := utils.CreateJWT(c, *u.Username, *u.ID)
	if err != nil {
		errHelper(c, "Error creating JWT Token", err)
		return err
	}

	user := UserCredentials {Username: u.Username, ID: u.ID}

	return utils.Tmpl(c, user, "view/home.html")
}

func AccountLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	u, err := UsersLogin(username, password)
	if err != nil {
		errHelper(c, "Error Login", err)
		return err
	}
	
	return handleNextScreenLogin(c, u)
}

func AccountCreate(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	u, err := UsersCreate(username, password)
	if err != nil {
		errHelper(c, "Error inserting user", err)
		return err
	}
	
	return handleNextScreenLogin(c, u)
}