package main

import (
    "./handler"
    "./template"
	"github.com/labstack/echo"
)
func main() {
    e := echo.New()
    e.Renderer = template.NewTemplates("./views/*.html")
    e.GET("/",handler.TimeLinePage())
	e.Logger.Fatal(e.Start(":1323"))
}
