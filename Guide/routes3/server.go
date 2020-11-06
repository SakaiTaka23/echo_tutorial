package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// グループ
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "pass" {
			// admin
			return true, c.String(http.StatusOK, "BasicAuthOK")
		}
		return false, nil
	}))

	// admin/get
	g.GET("/get", func(c echo.Context) error {
		return c.String(http.StatusOK, "/admingroup")
	})

	e.GET("/get", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/admin/get")
	})

	// admin/
	// g.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "/admin/ slash")
	// })

	e.Logger.Fatal(e.Start(":8000"))
}
