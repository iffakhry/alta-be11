package routes

import (
	"be11/apimvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	//inisialisasi echo
	e := echo.New()

	// set route endpoint
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)

	return e
}
