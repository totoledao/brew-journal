package handler

import (
	"log"

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

type ParamData struct {
	Id string
}

func Home(c echo.Context) error {
	data := TodoPageData{
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	err := utils.Tmpl(c, data, "view/components/layout.html", "view/home.html")
	if err != nil {
		log.Println(err)
	}

	return err
}

func Login(c echo.Context) error {
	err := utils.Tmpl(c, nil, "view/components/layout.html", "view/login.html")
	if err != nil {
		log.Println(err)
	}

	return err
}

func Show(c echo.Context) error {
	id := c.Param("id")
	data := ParamData{
		Id: id,
	}

	err := utils.Tmpl(c, data, "view/components/layout.html", "view/show.html")
	if err != nil {
		log.Println(err)
	}			

	return err
}