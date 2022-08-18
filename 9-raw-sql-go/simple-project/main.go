package main

// go get -u github.com/go-sql-driver/mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Guru struct {
	ID     int
	Nama   string
	Nip    string
	Status string
}

func main() {

	// membuat koneksi dari app golang ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/db_be11")

	defer db.Close()

	if err != nil {
		log.Fatal("error", err.Error())
	} else {
		fmt.Println("success connect to DB")
	}

	results, errselect := db.Query("SELECT id, nama, nip, status FROM guru")
	if errselect != nil {
		fmt.Println("error select", errselect.Error())
	}

	var dataGuruAll []Guru //penampung semua data guru
	for results.Next() {   // membaca per baris
		var rowguru Guru // penampung tiap baris
		errScan := results.Scan(&rowguru.ID, &rowguru.Nama, &rowguru.Nip, &rowguru.Status)
		if errScan != nil {
			fmt.Println("error scan", errScan.Error())
		}
		fmt.Println("row", rowguru)
		dataGuruAll = append(dataGuruAll, rowguru) //menambahkan ke slice
	}
	// fmt.Println(dataGuruAll)
	for _, v := range dataGuruAll {
		fmt.Println("ID:", v.ID, "Nama:", v.Nama, "NIP:", v.Nip, "Status:", v.Status)
	}

}
