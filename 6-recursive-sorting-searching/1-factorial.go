package main

import "fmt"

/*
6! = 6*5*4*3*2*1
5! = 5*4*3*2*1 =120
4! = 4*3*2*1
3! = 3*2*1
2! = 2*1
1! = 1

5! = 5*4!
4! = 4*3!
3! = 3*2!
2! = 2*1!
1! = 1

*/

func factorialLoop(angka int) int {
	var result int = 1
	for i := angka; i >= 1; i-- {
		result = result * i
	}
	return result
}

func factorialRecursive(angka int) int {
	//base case
	if angka == 1 {
		return 1
	} else { // recurrent relations
		// 5 * f(4) --> 5*4!
		// 4 * f(3) -> 4*6=24
		// 3 * f(2) -> 3*2=6
		// 2 * f(1) =2*1=2
		// 1 =1
		return angka * factorialRecursive(angka-1)
	}
}

func main() {
	fmt.Println(factorialLoop(5))      //120
	fmt.Println(factorialRecursive(5)) //120
}
