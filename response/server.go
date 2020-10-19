package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	e := echo.New()

	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	e.GET("/json", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/json-stream", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/json-pretty", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.JSONPretty(http.StatusOK, u, "  ")
	})

	e.GET("/json-blob", func(c echo.Context) error {
		encodedJSON := []byte{} // Encoded JSON from external source
		return c.JSONBlob(http.StatusOK, encodedJSON)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
