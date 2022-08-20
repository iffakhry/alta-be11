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
	db, err := sql.Open("mysql", "root:qwerty13@tcp(localhost:3306)/db_be11")

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
