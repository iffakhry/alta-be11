package main

import "fmt"

func main() {
	var age uint = 10
	fmt.Println(age)
	var data1 bool = true
	fmt.Println(data1)
	var data2 string = "alta"
	fmt.Println(data2)
	var data3 float64 = 3.5
	fmt.Println(data3)
	var data4 int = -1
	fmt.Println(data4)

	age = 40
	fmt.Println(age)

	var firstname, lastname string = "alterra", "academy"
	fmt.Println(firstname)
	fmt.Println(lastname)

	gender := "l"
	fmt.Println(gender)

	nilai := 100
	fmt.Println(nilai)
	nilai = 1000

	fmt.Println(gender, nilai, firstname)
	fmt.Printf("%T dan %T \n", firstname, lastname)

	const angka int = 10
	fmt.Println(angka)
	// angka = 50
	const nama string = "alta"
	fmt.Println(nama)

}
