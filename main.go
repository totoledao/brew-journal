package main

import (
	"github.com/labstack/echo/v4"
	"github.com/totoledao/brew-journal/handler"
	"github.com/totoledao/brew-journal/handler/utils"
)

func main() {
  // Echo instance
  app := echo.New()
	port := ":3000"

  // Serving files
  app.Static("/static", "static")

  // SQLite
  handler.InitDB()

  // Routes
  app.GET("/login", handler.Login)
  
  // Protected Routes
  app.GET("/", handler.Home, utils.JWTMiddleware())
	app.GET("/show/:id", handler.Show, utils.JWTMiddleware())

  // API
  app.POST("/api/account/login", handler.AccountLogin)
  app.POST("/api/account/create", handler.AccountCreate)
	
	// Live reload socket
	app.GET("/ping", utils.PingHandler)

  // Start server
  app.Logger.Fatal(app.Start(port))
}