package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	// membuat koneksi dari app golang ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/db_be11")

	if err != nil {
		log.Fatal("error", err.Error())
	} else {
		fmt.Println("success connect to DB")
	}

	return db

}
