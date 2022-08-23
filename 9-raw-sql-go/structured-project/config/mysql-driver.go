package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	// membuat koneksi dari app golang ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	// db, err := sql.Open("mysql", "root:qwerty123@tcp(localhost:3306)/db_be11")
	dbConnection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", dbConnection)

	if err != nil {
		log.Fatal("error sql Open ", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
		// panic("error connect db")
	} else {
		fmt.Println("success connect to DB")
	}

	return db

}
