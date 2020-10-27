package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))

	e.GET("/old", func(c echo.Context) error {
		return c.String(http.StatusOK, "old")
	})

	e.GET("/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "new")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
