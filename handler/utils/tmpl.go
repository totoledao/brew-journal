package utils

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

func Tmpl(c echo.Context, data any, src ...string) error {
	t := template.Must(template.ParseFiles(src...))
	return t.Execute(c.Response().Writer, data)
}