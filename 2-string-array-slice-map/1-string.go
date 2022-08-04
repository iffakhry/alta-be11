package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. len string
	sentence := "alterra"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// 2. compare string
	str1 := "abc"
	str2 := "abd"
	fmt.Println(str1 == str2)

	// 3. Contains
	var str = "something"
	var substr = "thi"
	res := strings.Contains(str, substr)
	fmt.Println(res) // true

	// 4. substring
	value := "cat;dog"
	// Take substring from index 4 to length of string.
	substring := value[4:len(value)]
	fmt.Println(substring)

	// 5. Replace
	s := "this[things]I would like to remove"
	t := strings.Replace(s, "[", "-", -1)
	fmt.Printf("%s\n", t)

	// 6. Insert
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)
}
