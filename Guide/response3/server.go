package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/file", func(c echo.Context) error {
		return c.File("test.png")
	})

	e.GET("/file-a", func(c echo.Context) error {
		return c.Attachment("test.png", "file_name.png")
	})

	e.GET("/file-i", func(c echo.Context) error {
		return c.Inline("test.png", "file-name.png")
	})

	e.GET("/blob", func(c echo.Context) (err error) {
		data := []byte(`0306703,0035866,NO_ACTION,06/19/2006
	  0086003,"0005866",UPDATED,06/19/2006`)
		return c.Blob(http.StatusOK, "text/csv", data)
	})

	e.GET("/stream", func(c echo.Context) error {
		f, err := os.Open("test.png")
		if err != nil {
			return err
		}
		return c.Stream(http.StatusOK, "image/png", f)
	})

	e.GET("/none", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/redirected")
	})

	e.GET("/redirected", func(c echo.Context) error {
		return c.String(http.StatusOK, "redirected")
	})

	e.GET("/hook", func(c echo.Context) error {
		c.Response().Before(func() {
			println("before response")
		})
		c.Response().After(func() {
			println("after response")
		})
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
