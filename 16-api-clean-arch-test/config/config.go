package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	SERVER_PORT int
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal(err)
	// }
	serverPortConv, errConv := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if errConv != nil {
		log.Fatal("error parse DB PORT")
		return nil
	}
	defaultConfig.SERVER_PORT = serverPortConv
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	dbPortConv, errConv := strconv.Atoi(os.Getenv("DB_PORT"))
	if errConv != nil {
		log.Fatal("error parse DB PORT")
		return nil
	}
	defaultConfig.DB_PORT = dbPortConv

	return &defaultConfig
}
