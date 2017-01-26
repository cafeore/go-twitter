package main

import (
    "./handler"
    "./template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
func main() {
    e := echo.New()
    e.Renderer = template.NewTemplates("./views/*.html")
    e.Use(middleware.Logger())
	e.Use(middleware.Recover())
    e.GET("/timeline",handler.TimeLinePage())
	e.Start(":1323")
}
