package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	fmt.Println("addr", (*e).name)
	fmt.Println("addr", e.name)
	(*e).name = newName
}

func (e *Employee) increaseSalary(newSalary int) {
	(*e).salary = newSalary
}

func main() {
	e := Employee{
		name:   "Ross Geller",
		salary: 1200,
	}
	fmt.Println("address", &e)
	fmt.Println("address", &e.name)

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")
	ep.increaseSalary(10000)
	// e after name change
	fmt.Println("e after name change =", e)
}
