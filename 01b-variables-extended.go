package main

import "fmt"

func main() {
	// Print with format specifier //
	var newName = "DonMax"
	var newAge = 25
	var newHeight = 1.83
	fmt.Printf("%T\n", newName)
	fmt.Printf("%T\n", newAge)
	fmt.Printf("%s is %d years old.\n", newName, newAge)
	fmt.Printf("%T %T %T\n", newName, newAge, newHeight)
}
