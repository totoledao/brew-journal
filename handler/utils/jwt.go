package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Username  string `json:"username"`
	ID string   `json:"id"`
	jwt.RegisteredClaims
}

const secret = "SUPERSECRET"

func CreateJWT(c echo.Context, username string, id string) error {
	cookieExpiration := time.Now().Add(time.Hour * 24 * 30) //30 days

	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(cookieExpiration), 
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error creating JWT Token")
		return err
	}

	// Set the JWT token as a cookie
	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = t
	cookie.Expires = cookieExpiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	// Send Token as response
	// c.JSON(http.StatusOK, echo.Map{		"token": t,	})
	return nil
}

func JWTMiddleware() echo.MiddlewareFunc{
	// Configure middleware with the custom claims type
  config := echojwt.Config{
    NewClaimsFunc: func(c echo.Context) jwt.Claims {
      return new(jwtCustomClaims)
    },
    SigningKey: []byte(secret),
  }
	
	// Return middleware to be used
	return echojwt.WithConfig(config)
}
