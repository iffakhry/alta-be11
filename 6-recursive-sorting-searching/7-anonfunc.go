package main

import "fmt"

func main() {
	value := func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}
	value()

	func() {

		fmt.Println("Welcome! to GeeksforGeeks guys")
	}()

	func(ele string) {
		fmt.Println(ele)
	}("GeeksforGeeks")
}
