package main

import (
	"fmt"
)

type Calculate interface {
	large() int
	keliling() int
	CetakHalo(str string, str1 string) (string, int)
}

type Square struct {
	side int
}

type Lingkaran struct {
	jari int
}

func (s Square) large() int {
	return s.side * s.side
}

func (s Square) keliling() int {
	return 4 * s.side
}

func (s Square) CetakHalo(data string, data1 string) (string, int) {
	return "halo" + data + data1, 1
}

func (s Square) Volume() int {
	return 0
}

func (l Lingkaran) large() int {
	return 3 * l.jari * l.jari
}

func (l Lingkaran) keliling() int {
	return 4 * l.jari
}

func (l Lingkaran) CetakHalo(data string, data1 string) (string, int) {
	return "halo", 1
}

func main() {
	var dimResult Calculate
	dimResult = Square{10}
	var hasil = dimResult.large() //100
	fmt.Println("hasil:", hasil)
	var hasilKeliling = dimResult.keliling()
	fmt.Println("hasil:", hasilKeliling)
	data1, data2 := dimResult.CetakHalo("hai", "hello")
	fmt.Println(data1, data2)

	var dimResultLingkaran Calculate
	dimResultLingkaran = Lingkaran{10}
	var hasilLingkaran = dimResultLingkaran.large() //300
	fmt.Println("hasil:", hasilLingkaran)
}
