package main

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	casbin_mw "github.com/labstack/echo-contrib/casbin"
)

func main() {
	e := echo.New()

	enforcer, err := casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
	e.Use(casbin_mw.Middleware(enforcer))

	e.Logger.Fatal(e.Start(":8000"))
}
