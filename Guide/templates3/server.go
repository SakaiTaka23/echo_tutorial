package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
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

func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}

func main() {
	e := echo.New()
	renderer := &TemplateRenderer{
		// templates: template.Must(template.ParseGlob("views/**/*")),
		templates: ParseTemplates(),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "test", map[string]interface{}{
			"name": "test",
		})
	})

	e.GET("/admin", func(c echo.Context) error {
		return c.Render(http.StatusOK, "adminindex", map[string]interface{}{
			"name": "admin",
		})
	})

	e.GET("/user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "userindex", map[string]interface{}{
			"name": "user",
		})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
