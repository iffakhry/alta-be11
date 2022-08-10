package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"cb", "ca", "bc"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	keys := []int{3, 2, 8, 1}
	// sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	// fmt.Println(keys)

	/*
		keys[i] < keys[j] --> untuk ascending
		keys[i] > keys[j] --> untuk descending
	*/
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	fmt.Println(keys)

	str := []string{"cb", "ca", "bc"}
	sort.SliceStable(str, func(i, j int) bool {
		return str[i] > str[j]
	})
	fmt.Println(str)

	// data := []int{3, 2, 8, 1}
	// for i := 0; i < len(data); i++ {
	// 	for j := i + 1; j < len(data); j++ {
	// 		fmt.Println("datai:", data[i], "dataj:", data[j])
	// 		if data[i] > data[j] {
	// 			fmt.Println("tukarr!!!")
	// 		} else {
	// 			fmt.Println("tetap")
	// 		}
	// 	}
	// }

}
