package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}

// // get user by id
// func GetUserController(c echo.Context) error {
// 	// your solution here
// }

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	// pertama := users[:2]
	// kedua := users[]
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	// id, _ := strconv.Atoi(c.Param("id"))
	// userUpdate := User{}
	// c.Bind(&userUpdate)
	// if userUpdate.Name != "" {
	// 	users[id].Name = userUpdate.Name
	// }
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"data":    user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
