package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		//swap
		elements[k], elements[minimal] = elements[minimal], elements[k]
	}
	return elements
}

func main() {
	var data = []int{2, 6, 1, 7, 9, 10, 17, 13, 4}
	fmt.Println("sebelum", data)
	result := selectionSort(data)
	fmt.Println("sesudah", result)

	a := "teh"
	b := "kopi"
	c := "jus"
	fmt.Println("a", a, "b", b, "c", c)
	a, b, c = c, a, b
	// var c = ""
	// c = a
	// a = b
	// b = c

	// bil := 1545

	fmt.Println("a", a, "b", b, "c", c)
}
