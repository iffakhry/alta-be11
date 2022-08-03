package main

import "fmt"

func main() {
	var suhu = 101
	var masker = true
	// if suhu == 36 {
	// 	fmt.Println("suhu normal", suhu)
	// }
	// if suhu < 40 {
	// 	fmt.Println("Anda demam")
	// }

	// if suhu == 36 {
	// 	fmt.Println("suhu normal", suhu)
	// } else {
	// 	fmt.Println("Anda demam")
	// }

	if suhu <= 36 && masker == true {
		fmt.Println("normal")
	} else if suhu > 37 && suhu < 40 {
		fmt.Println("panas")
	} else {
		fmt.Println("waspada")
	}

	if suhu == 35 {
		fmt.Println("suhu 35")
	} else if suhu > 100 {
		fmt.Println("else if")
	} else {
		fmt.Println("qqqqqq")
	}

}
