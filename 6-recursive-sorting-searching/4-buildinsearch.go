package main

import (
	"fmt"
	"sort"
)

func main() {
	elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	value := 10
	findIndex := sort.SearchInts(elements, value)
	fmt.Println("index", findIndex)
	if value == elements[findIndex] {
		fmt.Println("value ditemukan")
	} else {
		fmt.Println("value tidak ditemukan")
	}
}
