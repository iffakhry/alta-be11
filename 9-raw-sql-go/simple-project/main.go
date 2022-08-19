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

	// defer fmt.Println("halo 1")
	// defer fmt.Println("halo 2")
	// fmt.Println("halo 3")
	// defer fmt.Println("halo 4")

	// membuat koneksi dari app golang ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/db_be11")

	if err != nil {
		log.Fatal("error", err.Error())
	} else {
		fmt.Println("success connect to DB")
	}

	defer db.Close()

	fmt.Print("MENU:\n1. Baca Data\n2. Tambah Data\n")
	fmt.Println("Masukkan piihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
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
	case 2:
		{
			newGuru := Guru{}
			fmt.Println("Input Nama Guru:")
			fmt.Scanln(&newGuru.Nama)
			fmt.Println("Input NIP:")
			fmt.Scanln(&newGuru.Nip)
			fmt.Println("Input Status:")
			fmt.Scanln(&newGuru.Status)

			// sqlquery := fmt.Sprintf("insert into guru (nama, nip, status) values ('%s', '%s', '%s')", newGuru.Nama, newGuru.Nip, newGuru.Status)
			var query = "insert into guru (nama, nip, status) values (?, ?, ?)"
			statement, errPrepare := db.Prepare(query)
			if errPrepare != nil {
				fmt.Println("error prepare insert", errPrepare.Error())
			}

			result, errExec := statement.Exec(newGuru.Nama, newGuru.Nip, newGuru.Status)
			if errExec != nil {
				fmt.Println("error exec insert", errExec.Error())
			} else {
				// fmt.Println("sukses insert")
				row, errRow := result.RowsAffected()
				if errRow != nil {
					fmt.Println("error row insert", errRow.Error())
				}
				if row > 0 {
					fmt.Println("Success Insert Data")
				} else {
					fmt.Println("gagal insert")
				}
			}

		}
	case 3:
		{
			//update
		}

	case 4:
		{
			//delete
		}
	case 5:
		{
			//mendapatkan data guru berdasarkan id tertentu
		}

	}

}
