package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.BodyLimit("2M"))

	e.GET("/limit", func(c echo.Context) error {
		return c.String(http.StatusOK, "limit")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
