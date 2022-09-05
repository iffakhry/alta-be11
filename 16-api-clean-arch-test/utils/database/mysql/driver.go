package mysql

import (
	"be11/apiclean/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	// connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_be11?charset=utf8&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	// var err error
	// var db *gorm.DB
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to DB", err)
	}
	return db
}
