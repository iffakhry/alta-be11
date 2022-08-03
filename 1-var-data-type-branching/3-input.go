package main

import (
	"fmt"
)

// var (
// 	firstName, lastName, s string
// 	i                      int
// 	f                      float32
// 	input                  = "56.12 / 5212 / Go"
// 	format                 = "%f / %d / %s"
// )

func main() {
	// fmt.Println("Please enter your full name: ")
	// fmt.Scanln(&firstName, &lastName)
	// // fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)

	var nama string //fakhry
	fmt.Println("input nama:")
	fmt.Scanln(&nama)
	fmt.Println("Halo", nama)

	// in := bufio.NewReader(os.Stdin)

	// line, _ := in.ReadString('\n')
	// fmt.Println("halo", line)
}
