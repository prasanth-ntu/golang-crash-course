package main

import "fmt"

func main() {

	// animal := "Cat"
	animal := "Rat"

	switch animal {
	case "Cat":
		fmt.Println("Meow")
	case "Dog":
		fmt.Println("Woof")
	case "Frog":
		fmt.Println("Ribbit")
	case "Horse":
		fmt.Println("Neigh")
	default:
		fmt.Println("Don't know the animal:", animal)
	}

}
