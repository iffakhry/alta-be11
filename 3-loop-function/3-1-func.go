package main

import (
	"fmt"
	"math"
)

func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}

	if hour == 8 {
		fmt.Println("Semangat Kerja")
	}
}

func greetingReturn1(hour int) string {
	// var temp string
	if hour < 12 {
		return "Selamat Pagi"
	} else if hour < 18 {
		return "Selamat Sore"
	} else {
		return "Selamat Malam"
	}
	hour = hour + 1
	// fmt.Println("halo")
	return "halo"

}

func greetingReturn(hour int) string {
	var temp string
	if hour < 12 {
		temp = "Selamat Pagi"
	} else if hour < 18 {
		temp = "Selamat Sore"
	} else {
		temp = "Selamat Malam"
	}

	if hour == 8 {
		temp += "\nSemangat Kerja"
	}
	return temp
}

// multiple return value
func calculateCircle(diameter float64) (float64, float64, string) {
	var luas = math.Pi * math.Pow(diameter/2, 2) // 3.14 * r2
	var keliling = math.Pi * diameter            // 3.14 * d
	// return 2 value
	return keliling, luas, "sukses"
}

func main() {
	// greeting(20)
	// greeting(8)

	fmt.Println(greetingReturn(8))

	var data float64 = 10.0
	tempk, templ, temps := calculateCircle(data)
	// tempk, _, temps := calculateCircle(data)
	fmt.Println("keliling:", tempk, "luas:", templ, "str:", temps)
}
