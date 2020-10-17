package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	g := e.Group("/basic")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, c.String(http.StatusOK, "basic auth ok")
		}
		return false, c.String(http.StatusUnauthorized, "ng")
	}))

	e.Logger.Fatal(e.Start(":8000"))
}
