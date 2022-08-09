package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var data interface{}
	data = "alta"
	// fmt.Println(data)
	describe(data)
	data = 10
	describe(data)
	data = true
	describe(data)

	var data1 = map[string]interface{}{}
	data1["nama"] = "john"
	data1["umur"] = 20
	data1["isAktif"] = true
	fmt.Println(data1)
}
