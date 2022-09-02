package main

import (
	"be11/apiclean/config"
	"be11/apiclean/factory"
	"be11/apiclean/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	e := echo.New()
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
