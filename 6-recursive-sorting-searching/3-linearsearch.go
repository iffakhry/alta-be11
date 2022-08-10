package main

import "fmt"

func linierSearch(elements []int, x int) int {
	var count int
	for i := 0; i < len(elements); i++ {
		count++
		fmt.Println("count", count)
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var data = []int{4, 1, 6, 8, 9, 10, 30, 1, 90}
	// fmt.Println(linierSearch(data, 90))
	result := linierSearch(data, 11)
	if result != -1 {
		fmt.Println("data ditemukan di index", result)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
