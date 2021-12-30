package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
