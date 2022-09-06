package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", HelloWorldController)
	e.GET("/users", HelloUserController)
	e.GET("/books", HelloWorldController)
	e.Logger.Fatal(e.Start(":8000"))
}

func HelloWorldController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello world",
	})
}

func HelloUserController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello User",
	})
}
