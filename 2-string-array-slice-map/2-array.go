package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 15 --> 0 ... 14 --> len(variabel) = 15
	// 100 --> 0 ... 99 --> len(variabel) = 100

	var bilangan [10]int //0,1,2 .... 9
	fmt.Println(reflect.ValueOf(bilangan).Kind())
	bilangan[0] = 10
	// bilangan[1] = 20
	bilangan[2] = 15
	bilangan[2] = 20
	// bilangan[3] = 14
	bilangan[len(bilangan)-1] = 30 //bilangan[10-1] = 30
	fmt.Println(bilangan)
	fmt.Println(len(bilangan)) // length 10

	var country [3]string // 0,1,2
	country[0] = "indonesia"
	country[1] = "malaysia"
	country[2] = "singapore"
	// country[3] = "thailand"
	fmt.Println(len(country))
	fmt.Println(len(country[0]))

	country1 := [...]string{"indonesia", "malaysia"}
	fmt.Println(country1)

	// var data string = "hello"
	var data1 [3]int = [3]int{10, 20, 30}
	fmt.Println(data1)
	data2 := [3]int{10, 20, 30}
	fmt.Println(data2)

	var data3 = [...]int{14, 16, 12, 19}
	fmt.Println(len(data3))

	const index = 0
	var data4 [3]int = [3]int{index: 10, 2: 20}
	fmt.Println(data4)

	var data5 [2][3]string
	data5[0][0] = "indonesia"
	data5[0][1] = "malaysia"

}
