package main

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	e := echo.New()

	e.GET("/xml", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.XML(http.StatusOK, u)
	})

	e.GET("/xml-stream", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return xml.NewEncoder(c.Response()).Encode(u)
	})

	e.GET("/xml-pretty", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.XMLPretty(http.StatusOK, u, "  ")
	})

	e.GET("/xml-blob", func(c echo.Context) error {
		encodedXML := []byte{} // Encoded XML from external source
		return c.XMLBlob(http.StatusOK, encodedXML)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
