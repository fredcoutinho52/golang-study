package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// caution (convert to ASCII table)
	fmt.Println("Test " + string(123))

	// integer to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string to integer
	num, _ := strconv.Atoi("123") // using underline to ignore error
	fmt.Println(num - 122)

	// converting to boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("true")
	}
}
