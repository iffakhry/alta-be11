package main

import (
	"fmt"
	"math"
)

/*
bil prima adalah bilangan yang hanya habis dibagi oleh 1 dan bilangan itu sendiri

*/

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeBiasa(number int) bool {
	if number < 2 {
		return false
	}
	var temp int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			temp++
		}
	}
	if temp > 2 {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkPrime(1000000007)) //true
	// fmt.Println(checkPrimeBiasa(1000000007)) //true
}
