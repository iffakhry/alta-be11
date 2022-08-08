package main

import (
	calc "be11/unittest/calculator"
	"fmt"
)

/*
pertama, jalankan perintah go mod init, untuk generate file go.mod:

go mod init namaproject

NB: namaproject bisa bebas
*/

func Add(bil1, bil2 int) int {
	hasil := bil1 + bil2
	return hasil
}

func main() {
	hasil := Add(1, 2)
	fmt.Println(hasil)
	fmt.Println(calc.Tambah(5, 10))
	fmt.Println(calc.Kurang(10, 4))
}
