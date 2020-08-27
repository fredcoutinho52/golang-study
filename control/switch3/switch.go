package main

import (
	"fmt"
	"time"
)

func typeOf(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "Invalid type"
	}
}

func main() {
	fmt.Println(typeOf(2.3))
	fmt.Println(typeOf(1))
	fmt.Println(typeOf("Hi"))
	fmt.Println(typeOf(func() {}))
	fmt.Println(typeOf(time.Now()))
}
