package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	xs := fmt.Sprint(x) // retorna uma string
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x) // arredonda para duas casas decimais

	a := 1
	b := 1.999
	c := false
	d := "opa"
	fmt.Print("\n")
	fmt.Printf("%d %f  %.1f %t %s", a, b, b, c, d)
	fmt.Print("\n")
	fmt.Printf("%v %v  %.1v %v %v", a, b, b, c, d) // v é genérico
}

// IMPORTANTE
// Usar aspas duplas
// Usar vírgula para concatenar tipos diferentes
