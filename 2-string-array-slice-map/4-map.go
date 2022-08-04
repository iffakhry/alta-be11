package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)
	salary["agustus"] = 10000000
	salary["september"] = 15000000
	salary["september"] = 20000000
	salary["oktober"] = 20000000
	fmt.Println(salary)
	fmt.Println(len(salary))
	fmt.Println(salary["september"])
	// salary["agustus"] = 0
	delete(salary, "agustus") // remove value by key
	fmt.Println(salary["agustus"])
	isimap, isExist := salary["oktober"]
	fmt.Println("value:", isimap, "isexist:", isExist)

	a, b := 1, 2
	fmt.Println(a, b)

	var bulan = map[string]bool{}
	bulan["januari"] = true
	bulan["februari"] = false
	bulan["maret"] = true
	fmt.Println(bulan)
	delete(bulan, "maret")
	fmt.Println(bulan["maret"])
	fmt.Println(bulan["februari"])

	v, exist := bulan["februari"]
	if exist == false {
		fmt.Println("data tidak ada")
	} else {
		fmt.Println("data ada. dengan nilai:", v)
	}
	fmt.Println("value:", v, "isexist:", exist)

	var absensi = map[int]string{}
	absensi[1] = "adi"
	absensi[2] = "budi"
	absensi[5] = "cindi"
	v1, exist1 := absensi[3]
	if !exist1 { //exist1 == false --> false == false
		fmt.Println("data tidak ada")
	} else {
		fmt.Println("data ada. dengan nilai:", v1)
	}

	var data1 = map[string][]int{}
	data1["januari"] = []int{1, 2, 3}

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)

	// using make
	var salary_c = make(map[string]int)
	// var salary_c = map[string]int{}
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)
}
