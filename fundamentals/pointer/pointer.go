package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	// address of i
	p = &i
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}

// IMPORTANT
// Golang don't have arithmetic with pointers
// Pointer is a memory reference
// nil = null
