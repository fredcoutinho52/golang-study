package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	// forma reduzida para declarar uma var (recomendada)
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// outras formas de declaração
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

// IMPORTANTE
// O uso de uma variável é obrigatório
// É possível atribuir labels nas importações dos pacotes
