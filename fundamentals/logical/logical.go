package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buyTV50 := work1 && work2
	buyTV32 := work1 != work2 // XOR
	buyIceCream := work1 || work2

	return buyTV50, buyTV32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("TV 50: %t, TV 32: %t, Ice Cream: %t", tv50, tv32, iceCream)
}

// IMPORTANT
// Golang don't have XOR
