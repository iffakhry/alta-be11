package main

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Wallets  []Wallet
}

type Wallet struct {
	gorm.Model
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
}

// type User struct {
// 	ID uint
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`

// }

func InitDB() {
	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/db_be11?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Wallet{})
}

func init() {
	fmt.Println("init berjalan")
	InitDB()
	InitialMigration()
}

// func FailedResponseHelper(msg string) map[string]interface{} {
// 	return map[string]interface{}{
// 		"status":  "failed",
// 		"message": msg,
// 	}
// }

func FailedResponseHelper(c echo.Context, code int, msg string) error {
	return c.JSON(code, map[string]interface{}{
		"status":  "failed",
		"message": msg,
	})
}

func SuccessResponseHelper(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func SuccessDataResponseHelper(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

func main() {
	fmt.Println("main berjalan")
	//inisialisasi echo
	e := echo.New()

	// set route endpoint
	e.GET("/users", GetUserController)
	e.POST("/users", CreateUserController)
	//start server
	e.Logger.Fatal(e.Start(":80"))
}

func GetUserController(c echo.Context) error {

	var data []User
	tx := DB.Find(&data)
	if tx.Error != nil {
		return FailedResponseHelper(c, http.StatusInternalServerError, "failed to get data")
	}

	return c.JSON(http.StatusOK, SuccessDataResponseHelper("success to get data", data))
}

func CreateUserController(c echo.Context) error {
	userData := User{}
	err := c.Bind(&userData)
	if err != nil {
		return FailedResponseHelper(c, http.StatusBadRequest, "failed to bind data")
	}

	tx := DB.Create(&userData)
	if tx.Error != nil {
		return FailedResponseHelper(c, http.StatusInternalServerError, "failed to create data")
	}

	if tx.RowsAffected == 0 {
		return FailedResponseHelper(c, http.StatusInternalServerError, "failed to create data")
	}

	return c.JSON(http.StatusCreated, SuccessResponseHelper("success insert data"))
}
