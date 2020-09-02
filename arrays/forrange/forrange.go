package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for i, n := range numbers {
		fmt.Printf("%d) %d\n", i+1, n)
	}

	for _, num := range numbers {
		fmt.Println(num)
	}
}
