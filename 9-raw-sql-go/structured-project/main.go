package main

import (
	"be11/structured/config"
	"be11/structured/controllers/guru"
	"be11/structured/entities"
	"fmt"
)

// go get -u github.com/go-sql-driver/mysql

func main() {

	db := config.ConnectToDB()

	defer db.Close()

	fmt.Print("MENU:\n1. Baca Data\n2. Tambah Data\n")
	fmt.Println("Masukkan piihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		{
			result, err := guru.GetDataGuruAll(db)
			if err != nil {
				fmt.Println("error membaca data dari database", err)
			} else {
				for _, v := range result {
					fmt.Println("ID:", v.ID, "Nama:", v.Nama, "NIP:", v.Nip, "Status:", v.Status)
				}
			}
			// results, errselect := db.Query("SELECT id, nama, nip, status FROM guru")
			// if errselect != nil {
			// 	fmt.Println("error select", errselect.Error())
			// }

			// var dataGuruAll []entities.Guru //penampung semua data guru
			// for results.Next() {            // membaca per baris
			// 	var rowguru entities.Guru // penampung tiap baris
			// 	errScan := results.Scan(&rowguru.ID, &rowguru.Nama, &rowguru.Nip, &rowguru.Status)
			// 	if errScan != nil {
			// 		fmt.Println("error scan", errScan.Error())
			// 	}
			// 	// fmt.Println("row", rowguru)
			// 	dataGuruAll = append(dataGuruAll, rowguru) //menambahkan ke slice
			// }
			// // fmt.Println(dataGuruAll)
			// for _, v := range dataGuruAll {
			// 	fmt.Println("ID:", v.ID, "Nama:", v.Nama, "NIP:", v.Nip, "Status:", v.Status)
			// }
		}
	case 2:
		{
			newGuru := entities.Guru{}
			fmt.Println("Input Nama Guru:")
			fmt.Scanln(&newGuru.Nama)
			fmt.Println("Input NIP:")
			fmt.Scanln(&newGuru.Nip)
			fmt.Println("Input Status:")
			fmt.Scanln(&newGuru.Status)

			rowAffect, err := guru.InsertDataGuru(db, newGuru)
			if err != nil {
				fmt.Println("error insert data", err)
			} else {
				if rowAffect == 0 {
					fmt.Println("gagal insert data. row affected = 0")
				} else {
					fmt.Println("insert sukses. row affected = ", rowAffect)
				}
			}

			// // sqlquery := fmt.Sprintf("insert into guru (nama, nip, status) values ('%s', '%s', '%s')", newGuru.Nama, newGuru.Nip, newGuru.Status)
			// var query = "insert into guru (nama, nip, status) values (?, ?, ?)"
			// statement, errPrepare := db.Prepare(query)
			// if errPrepare != nil {
			// 	fmt.Println("error prepare insert", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(newGuru.Nama, newGuru.Nip, newGuru.Status)
			// if errExec != nil {
			// 	fmt.Println("error exec insert", errExec.Error())
			// } else {
			// 	// fmt.Println("sukses insert")
			// 	row, errRow := result.RowsAffected()
			// 	if errRow != nil {
			// 		fmt.Println("error row insert", errRow.Error())
			// 	}
			// 	if row > 0 {
			// 		fmt.Println("Success Insert Data")
			// 	} else {
			// 		fmt.Println("gagal insert")
			// 	}
			// }

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
			//mendapatkan data guru berdasarkan id tertentu --> 1 baris
		}

	}
}
