package main

import "fmt"

func main() {
	var bil1 int
	bil1 = 15
	bil2 := 30

	var bil3 = bil1 + bil2
	bil2 = 2
	bil3 = bil1 + bil2
	// bil4 := 90
	// bil5 := bil2 + bil4
	fmt.Println(bil3)

	bil6 := float32(15) / float32(2) // 15.0 / 2.0 --> 7.5
	// temp := float64(bil1 / bil2)     // 7
	// var bil6 float64 = float64(temp) // 7.0
	fmt.Println(bil6)

	var kata1 = "alterra"
	var kata2 = "academy"
	var hasil = kata1 + kata2
	fmt.Println(hasil)

	var index = 8
	var index2 = 8
	// index++
	fmt.Println(index == index2)
}
