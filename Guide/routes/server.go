package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "create")
}

func findUser(c echo.Context) error {
	return c.String(http.StatusOK, "find")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete")
}

func main() {
	e := echo.New()
	// Routes
	e.GET("/hello", hello)
	e.Any("/any", func(c echo.Context) error {
		return c.String(http.StatusOK, "any")
	})
	e.Match([]string{"GET", "POST"}, "/match", findUser)

	e.POST("/users", createUser)
	e.GET("/users", findUser)
	e.PUT("/users", updateUser)
	e.DELETE("/users", deleteUser)

	e.Logger.Fatal(e.Start(":8000"))
}
