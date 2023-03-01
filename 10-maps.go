package main

import "fmt"

func main() {
	cart := make(map[string]int)

	// Wwite //
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
}
