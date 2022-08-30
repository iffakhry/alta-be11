package main

//go get -u github.com/labstack/echo/v4

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

var data = []Article{
	{1, "Judul 1", "Deskripsi 1"},
	{2, "Judul 2", "Deskripsi 2"},
	{3, "Judul 3", "Deskripsi 3"},
}

func main() {
	//inisialisasi echo
	e := echo.New()
	// set route endpoint
	e.GET("/hello", HelloController)
	e.POST("/hello", PostHelloController)
	e.GET("/articles", GetArticlesController)
	//start server
	e.Start(":8080")
}

func HelloController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello world",
		"data":    1,
	})
}

func PostHelloController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hello Post",
		"data":    2,
	})
}

func GetArticlesController(c echo.Context) error {
	// return c.JSON(http.StatusOK, data)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}
