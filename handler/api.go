package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReqError struct {
  Message  string `json:"message" xml:"message"`
  Error  string `json:"error" xml:"error"`
}

type UserCredentials struct  {
 Username    string
 Password string
}

func errHelper(c echo.Context, message string, err error){	
		fmt.Println(message, err)

		e := &ReqError{
			Message:  message,
			Error:  err.Error(),
		}
  	c.JSON(http.StatusInternalServerError, e)
}

func AccountLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	err := UsersLogin(username, password)
	if err != nil {
		errHelper(c, "Error Login", err)
		return err
	}
	
	return c.String(http.StatusOK, "Login successful")
}

func AccountCreate(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := UsersCreate(username, password)
	if err != nil {
		errHelper(c, "Error inserting user", err)
		return err
	}
	
	return c.String(http.StatusOK, "Account created")
}