package main

import "fmt"

type Person struct {
	Name   string
	Age    uint
	Height uint
	Weight uint
}

type Mobil struct {
	Jenis      string
	Merk       string
	Transmisi  string
	BahanBakar string
	CC         uint
	Tahun      uint
	Warna      string
}

func main() {
	var person1 Person
	person1.Name = "John"
	person1.Age = 20
	person1.Weight = 70
	person1.Height = 170

	var person2 Person
	person2.Name = "Doe"
	person2.Name = "Jono"
	fmt.Println("person1:", person1.Name)
	fmt.Println("person2:", person2.Name)

	var name string
	name = "john"
	name = "doe"
	fmt.Println(name)

	var mobil1 Mobil
	mobil1.Jenis = "SUV"
	mobil1.BahanBakar = "Listrik"
	mobil1.CC = 3000
	mobil1.Merk = "ESEMKA"
	mobil1.Tahun = 2024
	mobil1.Warna = "putih"
	fmt.Println("mobil:", mobil1)
	fmt.Println("mobil:", mobil1.BahanBakar)

	var mobil2 Mobil
	mobil2.Jenis = "sedan"
	fmt.Println("mobil2:", mobil2)

	//long declaration with value
	var person3 Person = Person{"Tono", 25, 180, 60}
	fmt.Println("person3", person3)

	var person4 Person = Person{
		Name:   "Toni",
		Height: 180,
		Weight: 70,
	}
	fmt.Println("person4", person4)
	//sort declaration
	person5 := Person{"Tono", 25, 180, 60}
	fmt.Println("person5", person5)

	fmt.Println("address:", &person5.Name)

	// var person6 = new(Person)
}
