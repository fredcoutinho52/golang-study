package main

import "fmt"

func gradeConcept(g float64) string {
	var grade = int(g)
	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade <= 8 && grade >= 7:
		return "B"
	case grade <= 6 && grade >= 5:
		return "C"
	case grade <= 4 && grade >= 3:
		return "D"
	case grade <= 2 && grade >= 0:
		return "E"
	default:
		return "Invalid grade"
	}
}

func main() {
	fmt.Println(gradeConcept(9.8))
	fmt.Println(gradeConcept(6.9))
	fmt.Println(gradeConcept(2.1))
}
