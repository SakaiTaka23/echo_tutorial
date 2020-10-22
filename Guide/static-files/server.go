package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Static("/static", "assets")

	e.File("/favicon.svg", "images/favicon.svg")

	e.File("/", "public/index.html")

	e.Logger.Fatal(e.Start(":8000"))
}
