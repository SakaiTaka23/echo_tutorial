// ---------------------------------------------------------------------
/*

   server.go

                       Jun/11/2018
*/
// ---------------------------------------------------------------------
package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

// ---------------------------------------------------------------------
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// ---------------------------------------------------------------------
// サイトで共通情報
type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo{
	"サイトのタイトル",
}

// ---------------------------------------------------------------------
func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()

	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "こんにちは!")
	})

	e.GET("/page1", func(c echo.Context) error {
		// テンプレートに渡す値

		data := struct {
			ServiceInfo
			Content_a string
			Content_b string
			Content_c string
			Content_d string
		}{
			ServiceInfo: serviceInfo,
			Content_a:   "雨が降っています。",
			Content_b:   "明日も雨でしょうか。",
			Content_c:   "台風が近づいています。",
			Content_d:   "Jun/11/2018",
		}
		return c.Render(http.StatusOK, "page1", data)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

// ---------------------------------------------------------------------
