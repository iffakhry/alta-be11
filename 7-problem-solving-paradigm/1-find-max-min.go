package main

import (
	"fmt"
	"math"
)

func findMaxMin(data []int) (max int, min int) {
	min = math.MaxInt64
	max = math.MinInt64
	// max = data[0]
	// min = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		if data[i] < min {
			min = data[i]
		}
	}
	return
}

func main() {
	data := []int{4, 7, 1, 10, 30, 6, 8}
	max, min := findMaxMin(data)
	fmt.Println("max", max, "min", min)
}
