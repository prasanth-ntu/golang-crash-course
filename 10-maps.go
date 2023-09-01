package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Array vs. Slice vs. Map
		- Array: Fixed size, Ordered, Accessible by index
			e.g., var a [5]int
		- Slice: Dynamic size, Ordered, Accessible by index
			e.g., var s []int
		- Map: Dynamic size, Unordered collection, Accessible by key as they are key-value pairs. Maps are the equivalent of Python's dictionary.
			e.g., m = make(map[string]int)
	*/
	cart := make(map[string]int)

	// Write //
	cart["lamp"] = 3
	cart["bowl"] = 1
	cart["laptop"] += 1 // Since we haven't created laptop key before, it will assign 0 as a default value and perform addition
	cart["book"] = 0

	fmt.Println(cart)

	// Read //
	fmt.Println(cart["lamp"])

	itemQty, found := cart["lamp"]
	fmt.Println("cart[\"lamp\"]:", itemQty, found)

	itemQty, found = cart["book"]
	fmt.Println("cart[\"book\"]:", itemQty, found)

	itemQty, found = cart["noodles"]
	fmt.Println("cart[\"noodles\"]:", itemQty, found)

	// Delete //
	delete(cart, "laptop")
	itemQty, found = cart["laptop"]
	fmt.Println("cart[\"laptop\"]:", itemQty, found)

	// Iterate over map items //
	for key, value := range cart {
		fmt.Println("Key:", key, "=>", "Element:", value)
	}

	// Define another map //
	emails := make(map[string]string)

	// assign key value pairs to email map
	emails["Santh"] = "santh@gmail.com"
	emails["Sean"] = "sean@gmail.com"
	for key, value := range emails {
		fmt.Printf("Key: %s => Value: %s\n", key, value)
	}

	// Declare map and add KV at one shot //
	emails2 := map[string]string{
		"Santh": "santh@gmail.com",
		"Sean":  "sean@gmail.com",
		"Shawn": "shawn@gmail.com",
	}
	for key, value := range emails2 {
		fmt.Printf("Key: %s => Value: %s\n", key, value)
	}

	// Check if a key exists in the map //
	_, ok := emails2["Santh"]
	fmt.Println("Is 'Santh' in the map?", ok)
	// If the value is not present, it will return the default value of the data type //
	value, ok := emails2["Santh2"]
	fmt.Println("Is 'Santh2' in the map?", ok, ". The corresponding value is", value, "and the type of 'value' is", reflect.TypeOf(value))

}
