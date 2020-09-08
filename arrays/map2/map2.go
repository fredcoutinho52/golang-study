package main

import "fmt"

func main() {
	funcAndMoney := map[string]float64{
		"Jhon":  11325.5,
		"Ann":   15456.7,
		"Peter": 1200.0,
	}

	funcAndMoney["Jake"] = 1350.0
	delete(funcAndMoney, "Nonexistent")

	for name, money := range funcAndMoney {
		fmt.Println(name, money)
	}
}
