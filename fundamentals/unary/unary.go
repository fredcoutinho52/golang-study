package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	fmt.Println(x)

	y--
	fmt.Println(y)
}

// IMPORTANT
// Golang only have postfix unary operator
// Golang don't have ternary operator
