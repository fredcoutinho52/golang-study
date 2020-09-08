package main

import "fmt"

func main() {
	approved := make(map[int]string)

	approved[12345678978] = "Maria"
	approved[46546546587] = "Pedro"
	approved[31231231369] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[12345678978])
	delete(approved, 31231231369)
	fmt.Println(approved[31231231369])
}

// IMPORTANT
// maps need to be initialized
