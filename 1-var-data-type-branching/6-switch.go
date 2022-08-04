package main

import "fmt"

func main() {
	// var today int = 10
	// rain := "yes"
	// switch today {
	// case 1:
	// 	fmt.Printf("Today is Monday")
	// case 2:
	// 	fmt.Printf("Today is Tuesday")
	// case 3:
	// 	fmt.Printf("Today is Wednesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	// switch {
	// case today == 1 && rain == "yes":
	// 	fmt.Printf("Today is Monday")
	// case today == 2:
	// 	fmt.Printf("Today is Tuesday")
	// default:
	// 	fmt.Printf("Invalid Date")
	// }

	// if today == 1 {
	// 	fmt.Printf("Today is Monday")
	// } else if today == 2 {
	// 	fmt.Printf("Today is Tuesday")
	// } else {
	// 	fmt.Printf("Invalid Date")
	// }

	value := 100
	switch value {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		// fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

}
