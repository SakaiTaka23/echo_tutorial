package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func csrf(c echo.Context) error {
	token := c.Get("csrf").(string)
	fmt.Print(token)
	return c.String(http.StatusOK, "OK")
}

func check(c echo.Context) error {
	token := c.Get("csrf").(string)
	name := c.Get("name").(string)
	fmt.Print(token)
	fmt.Print(name)
	return c.String(http.StatusOK, "OK")
}

func main() {
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("Request Body: %v\n", string(reqBody))
		fmt.Printf("Response Body: %v\n", string(resBody))
	}))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	e.GET("/csrf", csrf)
	e.GET("/form", func(c echo.Context) error {
		return c.Render(http.StatusOK, "form.html", map[string]interface{}{
			"csrf":c.Get("csrf"),
		})
	})
	e.POST("/form", check)

	e.Logger.Fatal(e.Start(":8000"))
}
