package main

import "fmt"

type Animal struct {
	animal_type string
	age         int
	gender      bool // false: female, true: male
	weight      float32
}

// Example of pointer receiver (*Animal) => Method can modify the struct instance
// Add birthday method to Animal struct to increment age by 1
func (a *Animal) birthday() {
	fmt.Println("Happy birthday", a.animal_type)
	a.age += 1
}

// Example of value receiver (Animal) => Method cannot modify the struct instance
// Print the age of the animal
func (a Animal) printAge() {
	fmt.Println("Age of the animal:", a.age)
}

func main() {
	fmt.Println("Intro to Structs")

	// Let's create a new Animal instance
	var pumba = Animal{"lion", 10, false, 200.5}
	pumba.birthday()
	pumba.printAge()
	fmt.Println(pumba)
}
