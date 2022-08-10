package main

import "fmt"

func Cetak(angka ...int) {
	for _, v := range angka {
		fmt.Println(v)
	}
}

func main() {
	a, b, c, d := 10, 20, 30, 40
	Cetak(a, b, c, d, 10, 20, 50)
	// var data = []int{1, 2, 3, 4}
	// fmt.Println(&data[0])
}
