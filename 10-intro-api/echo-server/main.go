package main

//go get -u github.com/labstack/echo/v4

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int
	Name     string
	Age      uint
	Password string
}

type Book struct {
	ID        int    `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

type Article struct {
	ID      int
	Title   string
	Content string
}

func main() {
	//inisialisasi echo
	e := echo.New()

	// set route endpoint
	e.GET("/hello", HelloController)
	e.POST("/hello", PostHelloController)
	e.GET("/articles", GetArticlesController)
	e.GET("/books", GetBooksController)
	e.GET("/books/:id_book", GetBookByIdController)

	e.POST("/users", CreteUserController)
	e.POST("/books", CreteBookController)
	//start server
	e.Logger.Fatal(e.Start(":8080"))
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
	title := c.QueryParam("title")
	angka := c.QueryParam("angka")
	var data = []Article{
		{1, "Judul 1", "Deskripsi 1"},
		{2, "Judul 2", "Deskripsi 2"},
		{3, "Judul 3", "Deskripsi 3"},
	}
	angkaconv, err := strconv.Atoi(angka)
	if angka != "" {
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "failed",
			})
		}
		angkaconv = angkaconv + 1
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
		"title":   title,
		"angka":   angkaconv,
	})
}

func GetBooksController(c echo.Context) error {
	// var bookdata = Book{
	// 	ID:        1,
	// 	Title:     "One Piece",
	// 	Author:    "Oda",
	// 	Publisher: "Gramedia",
	// }
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"data":    bookdata,
	// })

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "failed",
		"data":    nil,
	})
}

func GetBookByIdController(c echo.Context) error {
	id := c.Param("id_book")
	fmt.Println("id", id)
	idconv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to convert param",
		})
	}
	var bookdata = Book{
		ID:        idconv,
		Title:     "One Piece",
		Author:    "Oda",
		Publisher: "Gramedia",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    bookdata,
	})
}

func CreteUserController(c echo.Context) error {
	name := c.FormValue("name")
	age, _ := strconv.Atoi(c.FormValue("age"))
	password := c.FormValue("password")
	id, _ := strconv.Atoi(c.FormValue("id"))
	// file, _ := c.FormFile("images")

	userdata := User{
		ID:       id,
		Name:     name,
		Age:      uint(age),
		Password: password,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert data",
		"data":    userdata,
	})
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id":       id,
	// 	"name":     name,
	// 	"age":      age,
	// 	"password": password,
	// })
}

func CreteBookController(c echo.Context) error {
	bookdata := Book{}
	err := c.Bind(&bookdata)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert book",
		"data":    bookdata,
	})
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"id":        bookdata.ID,
	// 	"author":    bookdata.Author,
	// 	"title":     bookdata.Title,
	// 	"publisher": bookdata.Publisher,
	// })
}
