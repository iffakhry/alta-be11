package main

import "fmt"

func main() {
	var name string = "John"
	fmt.Println("name address", &name) // output memory address
	fmt.Println("name value", name)    //john

	var namePointer *string = &name
	fmt.Println("namePointer", namePointer)        // output memory address
	fmt.Println("namePointer value", *namePointer) //john

	fmt.Println("===============")
	name = "Doe"
	name = "Jono"
	*namePointer = "Tono"                          // mengganti value yg disimpan di alamat memory tsb dengan Tono
	fmt.Println("name address", &name)             // output memory address
	fmt.Println("name value", name)                //Tono
	fmt.Println("namePointer", namePointer)        // output memory address
	fmt.Println("namePointer value", *namePointer) //Tono
	fmt.Println("namePointer value", len(*namePointer))
	// var temp = *namePointer

	if *namePointer == "Tono" {
		fmt.Println("namanya tono")
	}
	number_a := 25
	var number_b *int
	if number_b == nil {
		fmt.Println("number_b is", number_b)
		number_b = &number_a
		fmt.Println("number_b after init : is", *number_b)
	}

}
