package main

import "fmt"

// pass by value
func ubahAngka(bil int) int {
	fmt.Println("angka awal:", bil)
	bil = 100
	return bil
}

//pass by reference
func ubahAngkaPointer(bil *int) int {
	fmt.Println("alamat angka awal:", bil)
	fmt.Println("angka awal:", *bil)
	*bil = 200
	return *bil
	// result := *bil + 500
	// return result
}

func main() {
	var nilai = 90
	fmt.Println("ubahAngka:", ubahAngka(nilai)) // 100
	fmt.Println("nilai:", nilai)                //90
	fmt.Println("=======")
	fmt.Println("ubahAngkaPointer:", ubahAngkaPointer(&nilai)) //200
	fmt.Println("nilai:", nilai)                               //90
	fmt.Println("alamat nilai:", &nilai)                       //90
}
