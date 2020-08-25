package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // type float64 inferred by the compiler

	// reduced form to declare a var (recommended)
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// other forms of declaration
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "oi"
	fmt.Println(g, h, i)
}

// IMPORTANT
// Use an declared variable is mandatory
// Is possible to atribute labels for imports
