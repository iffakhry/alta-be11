package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// 0,1,2,3,4
		if i == 1 { // false, true, false, false, false
			continue
		}
		if i > 3 { // false, _ , false, false, true
			break
		}
		fmt.Println(i) // 0, 2,3
	}

	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(data); i++ { // i < 8 --> 7
		fmt.Println("i:", i)
		fmt.Println(data[i])
	}

	for i := len(data) - 1; i >= 0; i-- {
		fmt.Println(data[i])
	}

	for index, value := range data {
		fmt.Println("data index:", data[index])
		fmt.Println("data value:", value)
	}

	// fmt.Println(data[0])
	// fmt.Println(data[1])
	// fmt.Println(data[2])
	// fmt.Println(data[3])
	// fmt.Println(data[4])
	// fmt.Println(data[5])
	// fmt.Println(data[6])
	// fmt.Println(data[7])

}
