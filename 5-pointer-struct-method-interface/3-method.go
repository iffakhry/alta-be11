package main

import "fmt"

type Employee struct {
	FirstName, LastName string
	Age                 uint
}

//function
func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

//method
func (e Employee) fullNameMethod() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) IncreaseAge(angka int) uint {
	result := e.Age + uint(angka)
	return result

}

func main() {
	emp1 := Employee{
		FirstName: "John",
		LastName:  "Doe",
		Age:       20,
	}
	var temp = 10

	fmt.Println(fullName(emp1.FirstName, emp1.LastName))
	fmt.Println(emp1.fullNameMethod())
	fmt.Println(emp1.IncreaseAge(temp))
}
