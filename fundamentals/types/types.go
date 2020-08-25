package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer numbers
	fmt.Println(1, 2, 100)
	fmt.Println("Integer is", reflect.TypeOf(32000))

	// unsigned (only positives)... uint8 uint16 uint32 uint64
	var b byte = 255 // byte -> uint8
	fmt.Println("byte is", reflect.TypeOf(b))

	// signed... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The int max value is", i1)
	fmt.Println("i1 type is", reflect.TypeOf(i1))

	var i2 rune = 'a' // represents unicode table (int32)
	fmt.Println(i2)
	fmt.Println("rune is", reflect.TypeOf(i2))

	// float numbers (float32, float64)
	var x float32 = 49.99
	fmt.Println("x type is", reflect.TypeOf(x))
	fmt.Println("49.99 type is", reflect.TypeOf(49.99)) // float64 by default

	// boolean type
	bo := true
	fmt.Println("bo type is", reflect.TypeOf(bo))
	fmt.Print(!bo)

	// string type
	s1 := "Hello, my name is Fred"
	fmt.Println(s1 + "!")
	fmt.Println("string size is", len(s1))

	// string with multiple lines
	s2 := `Hello
	my
	name
	is
	Fred`
	fmt.Println("string size is", len(s2))
}

// IMPORTANT
// Use double quotes for strings
// Use single quotes for characteres
