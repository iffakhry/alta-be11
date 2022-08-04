package main

import (
	"fmt"
	"reflect"
)

/*
1. buat perulangan dari 1 sampai n
2. cek apakah n % index == 0
3. jika ya, maka simpan index ke slice (pakai append)
*/

func main() {
	var bilangan = []int{1, 3, 5, 7, 9}
	fmt.Println(bilangan)
	fmt.Println(len(bilangan))

	// menambah data ke slice
	bilangan = append(bilangan, 10)
	fmt.Println("---------------")
	fmt.Println(bilangan)
	fmt.Println(len(bilangan))

	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [5]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var bagianprimes []int = primes[1:4]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(bagianprimes).Kind())
	fmt.Println(bagianprimes)

	// menggabungkan 2 slice
	var bilangan1 = []int{1, 2, 3}
	var bilangan2 = []int{4, 5, 6, 7}
	var bilangan3 []int
	var tampung []int
	// bilangan1 = append(bilangan1, bilangan2...)
	// bilangan3 = append(bilangan1, bilangan2...)
	bilangan3 = append(bilangan1[1:], bilangan2[1:]...)
	fmt.Println(bilangan3)
	fmt.Println(bilangan1[0] + bilangan1[1])

	var temp = bilangan1[0] + bilangan1[1]  //1+2
	tampung = append(tampung, temp)         //append ke tampung
	var temp1 = bilangan1[1] + bilangan1[2] // 2+3
	tampung = append(tampung, temp1)        //append ke tampung
	fmt.Println("tampung", tampung)

	var data1 = []bool{true, false, true}
	data1 = append(data1, true)
	fmt.Println(data1)

	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purple")

	copied_colors := make([]string, 4)
	// hanya bisa copy slice
	copy(copied_colors, colors) // copy from colors to copied_colors
	fmt.Println(copied_colors)

	// var primes1 = [5]int{2, 3, 5, 7, 11}
	// copied_primes1 := make([]int, 4)

	// copy(copied_primes1, primes1) // copy from colors to copied_colors
	// fmt.Println(copied_primes1)

}
