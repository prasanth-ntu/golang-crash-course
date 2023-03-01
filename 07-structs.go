package main

import "fmt"

func main() {
	type Animal struct {
		class  string
		age    int
		gender bool // false: Male, true: Female
	}

	// One way to define a struct variable //
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

	// Alternative way to define struct variable //
	var leo = Animal{"lion", 2, false} // Order matters
	fmt.Println("leo:", leo)

	// Another way to define struct variable //
	var nemo = Animal{} // Assigns default value to each variable in struct
	fmt.Println("nemo=", nemo)
	nemo.class = "Fish"
	fmt.Println("nemo=", nemo)

	// Another way to create a struct variable without "type" //
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
