package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := 10
	b := &a // Assigning b to pointer of a (memory address of where the value is stored)

	//fmt.Printf("a=%T, b=%T\n", a, b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	fmt.Println(a, b)

	// Read the value from memory address using *
	fmt.Println(a, *b)
	fmt.Println(a, &a, *&a)

	// Change value with pointer
	*b = 5
	fmt.Println(a, b)

	a = 15 // b still points to the same address
	fmt.Println(a, b)

}
