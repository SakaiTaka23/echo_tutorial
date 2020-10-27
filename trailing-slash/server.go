package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())

	e.GET("/slash", func(c echo.Context) error {
		return c.String(http.StatusOK, "/slash")
	})

	e.GET("/slash/", func(c echo.Context) error {
		return c.String(http.StatusOK, "/slash/")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
