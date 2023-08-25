package main

import "fmt"

func main() {
	/*
		=== APPROACH 1 ===
		Define a new structure "Animal".
		struct is a collection of fields with different types.
	*/
	type Animal struct {
		class  string
		age    int
		gender bool // false: Male, true: Female
	}

	/*
		=== Different ways to define struct variable
		(a.k.a Animal instance) ===
	*/
	// --- Approach 1 --- //
	var teddy = Animal{
		class:  "bear",
		age:    24,
		gender: true,
	} // Order does not matter

	fmt.Println(teddy)
	fmt.Println(teddy.class, teddy.age, teddy.gender)

	teddy.age = teddy.age + 1
	teddy.class = "cute bear"
	fmt.Println("teddy:", teddy)

	// --- Approach 2 --- //
	var leo = Animal{"lion", 2, false} // Order matters
	fmt.Println("leo:", leo)

	// --- Approach 3 --- //
	var nemo = Animal{} // Assigns default value to each variable in struct
	fmt.Println("nemo=", nemo)
	nemo.class = "Fish"
	fmt.Println("nemo=", nemo)

	/*
		=== APPROACH 2 ===
		Another way to create a struct variable without "type"
		i.e., Anonymous structure (a structure without a name)
	*/
	pumba := struct {
		class  string
		age    int
		gender bool
	}{
		class:  "lion",
		age:    10,
		gender: false,
	}
	fmt.Println("pumba=", pumba)
}
