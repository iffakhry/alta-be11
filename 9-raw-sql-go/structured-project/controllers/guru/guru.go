package guru

import (
	"be11/structured/entities"
	"database/sql"
)

func GetDataGuruAll(db *sql.DB) ([]entities.Guru, error) {

	results, errselect := db.Query("SELECT id, nama, nip, status FROM guru")
	if errselect != nil {
		// fmt.Println("error select", errselect.Error())
		return nil, errselect
	}

	var dataGuruAll []entities.Guru //penampung semua data guru
	for results.Next() {            // membaca per baris
		var rowguru entities.Guru // penampung tiap baris
		errScan := results.Scan(&rowguru.ID, &rowguru.Nama, &rowguru.Nip, &rowguru.Status)
		if errScan != nil {
			// fmt.Println("error scan", errScan.Error())
			return nil, errScan
		}
		// fmt.Println("row", rowguru)
		dataGuruAll = append(dataGuruAll, rowguru) //menambahkan ke slice
	}
	// fmt.Println(dataGuruAll)
	// for _, v := range dataGuruAll {
	// 	fmt.Println("ID:", v.ID, "Nama:", v.Nama, "NIP:", v.Nip, "Status:", v.Status)
	// }
	return dataGuruAll, nil
}

func InsertDataGuru(db *sql.DB, newGuru entities.Guru) (int, error) {
	// sqlquery := fmt.Sprintf("insert into guru (nama, nip, status) values ('%s', '%s', '%s')", newGuru.Nama, newGuru.Nip, newGuru.Status)
	var query = "insert into guru (nama, nip, status) values (?, ?, ?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		// fmt.Println("error prepare insert", errPrepare.Error())
		return -1, errPrepare
	}

	result, errExec := statement.Exec(newGuru.Nama, newGuru.Nip, newGuru.Status)
	if errExec != nil {
		// fmt.Println("error exec insert", errExec.Error())
		return -1, errExec
	} else {
		// fmt.Println("sukses insert")
		row, errRow := result.RowsAffected()
		if errRow != nil {
			// fmt.Println("error row insert", errRow.Error())
			return 0, nil
		}
		// if row > 0 {
		// 	fmt.Println("Success Insert Data")
		// } else {
		// 	fmt.Println("gagal insert")
		// }
		return int(row), nil
	}
}
