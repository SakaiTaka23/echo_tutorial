package main

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	g := e.Group("/login")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "/auth")
	})

	e.GET("/every", func(c echo.Context) error {
		return c.String(http.StatusOK, "/every")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
