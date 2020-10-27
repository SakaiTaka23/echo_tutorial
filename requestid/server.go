package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
        return next(c)
    }
}

func main() {
	e := echo.New()
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.Use(ServerHeader)

	e.Use(middleware.RequestID())

	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Printf("Request Body: %v\n", string(reqBody))
	// 	fmt.Printf("Response Body: %v\n", string(resBody))
	// }))

	e.GET("/", func(c echo.Context) error {
		output := fmt.Sprintf("%#v", c.Request().Header)
		fmt.Print(output)

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
