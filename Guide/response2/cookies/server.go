package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all the cookies")
}

func deleteCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	// cookie := new(http.Cookie)
	// cookie.Name = "username"
	// //cookie.Value = "jon"
	// cookie.Expires = time.Now()
	// c.SetCookie(cookie)
	return c.String(http.StatusOK, "delete a cookie")
}

func main() {
	e := echo.New()

	e.GET("/write", writeCookie)
	e.GET("/read", readCookie)
	e.GET("/readall", readAllCookies)
	e.GET("/delete", deleteCookie)

	e.Logger.Fatal(e.Start(":8000"))
}
