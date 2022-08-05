package main

import "fmt"

func CetakHello() { //PascalCase
	fmt.Println("Hello guys")
	fmt.Println("Perkenalkan")
	fmt.Println("Salam kenal ya")

}

func cetakApaKabar() { //camelCase
	fmt.Println("Apa Kabar")
}

func HitungLuasSegitiga() {
	//
	//
	//
}

//return string
func Sapa() string {
	var data = "Halo semuanya"
	return data
}

//return int
func ReturnInt() int {
	var i = 0
	i = i + 10
	return i
}

//func with param
func WithParam(data1 int) {
	fmt.Println("data1", data1)
	var temp = data1 + 10
	fmt.Println(temp)
}

func Param2(data1 int, data2 string) {
	fmt.Println(data1, data2)
}

func Param2Return(data1 int, data2 string) string {
	fmt.Println(data1, data2)
	return data2
}

func main() {
	fmt.Println("haii")
	CetakHello()
	CetakHello()
	CetakHello()
	cetakApaKabar()
	var datasapa = Sapa()
	fmt.Println(datasapa)
	var tempint = ReturnInt()
	fmt.Println(tempint)
	fmt.Println(ReturnInt())

	var angka = 100
	WithParam(angka)

	Param2(120, "abc")

	var tempparam = Param2Return(1, "ab")
	fmt.Println(tempparam)
}
