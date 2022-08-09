package main

import (
	"fmt"
	"strings"
)

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string ", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)
	}
}

func main() {
	var secret interface{}

	secret = 2 //int
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	var data interface{}
	data = 10
	explain(data)
	data = "alta"
	explain(data)
	explain(secret)
}
