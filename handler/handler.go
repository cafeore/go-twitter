package handler
import (
	"../database"
	"github.com/labstack/echo"
	"net/http"
)

func TimeLinePage() echo.HandlerFunc{
	return func(c echo.Context)error{
		return c.Render(http.StatusOK,"go",database.GetTweets())
	}
}