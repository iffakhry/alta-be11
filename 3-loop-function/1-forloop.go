package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	fmt.Println("hello ")
	// }

	// index := 0
	// for index < 5 { // 0,1,2,3,4,5
	// 	// 0,1,2,3,4
	// 	fmt.Println(index) // 0,1,2,3,4
	// 	fmt.Println("hello ")
	// 	index++ // 1,2,3,4,5
	// }

	index := 0
	for index < 5 { // 0,1,2,3,4,5
		// 0,1,2,3,4
		index++            // 1,2,3,4,5
		fmt.Println(index) // 1,2,3,4,5
		fmt.Println("hello ")
	}
}
