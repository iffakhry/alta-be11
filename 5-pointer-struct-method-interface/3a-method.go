package main

import "fmt"

type Person struct {
	name string // Both non exported fields.
	age  int
}

func (P Person) GetName() string {
	// P.name = "andi"
	return P.name + " amazing!"
}

func (P *Person) IncreaseAge() {
	P.age = P.age + 1
	// p.age = 51
}

func main() {
	PersonA := Person{"John", 50}
	fmt.Printf("%v\n", PersonA)
	fmt.Println(PersonA.GetName())
	fmt.Println(PersonA.name)

	PersonB := Person{"Doe", 40}
	fmt.Printf("%v\n", PersonB)
	fmt.Println(PersonB.GetName())

	fmt.Println("sebelum:", PersonA.age)
	PersonA.IncreaseAge()
	fmt.Println("sesudah:", PersonA.age)
}
