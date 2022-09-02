package main

import (
	"be11/apimvc/config"
	"be11/apimvc/middlewares"
	"be11/apimvc/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	//start server
	e.Logger.Fatal(e.Start(":8000"))
}
