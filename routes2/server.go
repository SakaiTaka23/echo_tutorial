package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func users(c echo.Context) error {
	return c.String(http.StatusOK, "users")
}

func main() {
	e := echo.New()

	//パラメーター
	e.GET("/users", users)

	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})

	e.GET("/users/1/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})

	// グループ
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true,c.String(http.StatusOK, "basic auth ok")
		}
		return false, nil
	}))

	// ルート名
	route := e.POST("/name", func(c echo.Context) error {
		return c.String(http.StatusOK, "/name")
	})
	route.Name = "create-user"

	e.GET("/inline/name", func(c echo.Context) error {
		return c.String(http.StatusOK, "/inline/name")
	}).Name = "get-user"

	// ルート出力
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	_ = ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":8000"))
}
