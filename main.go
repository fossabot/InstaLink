package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/message/", func(c echo.Context) error {
		return c.String(http.StatusOK, "InstaLink is now under the construction.")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
