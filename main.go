package main

import (
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/totoledao/brew-journal/handler/utils"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func tmpl (c echo.Context, data any,  src ...string){
	t := template.Must(template.ParseFiles(src...))
	t.Execute(c.Response().Writer, data)
}

func main() {
  // Echo instance
  app := echo.New()
	port := ":3000"

  //Styles
  app.Static("/static", "static")

	

  // Routes
  app.GET("/", func(c echo.Context) error {
			data := TodoPageData{					
					Todos: []Todo{
							{Title: "Task 1", Done: false},
							{Title: "Task 2", Done: true},
							{Title: "Task 3", Done: true},
					},
			}
			tmpl(c, data, "view/index.html", "view/components/buttons/clickMe.html")
			return nil
	})

  app.POST("/clicked", func(c echo.Context) error {
		tmpl(c, nil, "view/components/buttons/clicked.html")
		return nil
	})
 
	app.GET("/ping", utils.PingHandler)

  // Start server
  app.Logger.Fatal(app.Start(port))
}