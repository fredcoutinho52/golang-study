package main

import "fmt"

func main() {
	funcByLetter := map[string]map[string]float64{
		"G": {
			"Gary":     15456.7,
			"Gulliver": 8456.7,
		},
		"J": {
			"Jhon": 11235.5,
		},
		"P": {
			"Peter": 1200.0,
		},
	}

	delete(funcByLetter, "P")

	for letter, funcs := range funcByLetter {
		fmt.Println(letter)

		for name, money := range funcs {
			fmt.Println(name, money)
		}
	}
}
