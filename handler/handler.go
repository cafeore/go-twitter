package handler
import (
	"../database"
	"github.com/labstack/echo"
	"net/http"
)
func LoginPage()echo.HandlerFunc{
	return func(c echo.Context)error{
		return c.Render(http.StatusOK,"login","")
	}
}

func TimeLinePage() echo.HandlerFunc{
	return func(c echo.Context)error{
		return c.Render(http.StatusOK,"go",database.GetTweets())
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("username")
       	pass := c.FormValue("password")
       	if database.GetUser(name,pass) {
			return c.Render(http.StatusOK,"go",database.GetTweets())
		}
		return c.Render(http.StatusOK,"fail","ログイン失敗")
	}
}
func CreateAccount() echo.HandlerFunc{
	return func(c echo.Context) error {
		name := c.FormValue("username")
       	pass := c.FormValue("password")
       	if database.GetUser(name,pass) {
			return c.Render(http.StatusOK,"fail","すでにそのＩＤは使われています")
		}
		return c.Render(http.StatusOK,"fail","アカウント作成完了しました"+"\n"+"戻ってログインしてください")
	}
}