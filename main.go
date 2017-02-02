package main

import (
    "./handler"
    "./template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
var user string
func main() {
    e := echo.New()
    e.Renderer = template.NewTemplates("./views/*.html")
    e.Use(middleware.Logger())
	e.Use(middleware.Recover())
    e.GET("/login",handler.LoginPage())
    e.GET("/timeline",handler.TimeLinePage())
    e.POST("/timeline",handler.Login())
    e.POST("/new",handler.CreateAccount())
    e.POST("/tweet",handler.Tweet())
	e.Start(":1323")
}
