package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %v\n", string(reqBody))
		fmt.Printf("Response Body: %v\n", string(resBody))
	}))

	e.GET("/dump", func(c echo.Context) error {
		return c.String(http.StatusOK, "dump")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
