package routes

import (
	"be11/apimvc/controllers"
	"be11/apimvc/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	//inisialisasi echo
	e := echo.New()

	// set route endpoint
	e.POST("/auth", controllers.LoginController)
	e.GET("/users", controllers.GetUserController, middlewares.JWTMiddleware())
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.POST("/wallets", controllers.CreateWalletController)
	e.GET("/wallets", controllers.GetWalletController)

	return e
}
