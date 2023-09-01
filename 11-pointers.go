package main

import "fmt"

func main() {
	// Standard //
	var a = 1
	var b = a // We get a copy of 'a' and assign it to 'b'
	fmt.Println(a, b)

	b += 1
	fmt.Println(a, b)

	// With pointers //
	var c = 1
	var d *int // d is a pointer
	d = &c     // d is pointing to c (i.e., memory address of c
	/*
		Explanation:
			- c: the value of c,
			- *d: value at the memory address that d points to,
			- d: memory address stored in d,
			- &c: memory address of c,
			- &d: memory address of d.
	*/
	fmt.Println(c, *d, d, &c, &d)

	*d += 1 // Since d is pointing to c, original variable stored in memory
	fmt.Println(c, *d, d, &c, &d)
}
