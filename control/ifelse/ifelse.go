package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved")
	} else {
		fmt.Println("Disapproved")
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
}
