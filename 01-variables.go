package main

import "fmt"

func main() {
	fmt.Println("Hello my friends")

	// Declaring and Assigning a value to variable //
	var favBook = "Harry Potter" // Data Type of this variable is inferred as a string
	fmt.Println(favBook)
	// Assigning (Updating) value of existing variable
	favBook = "Deep Work"
	fmt.Println(favBook)

	// Since Go is strongly typed language, we cannot assign a integer value to a variable that stotes string //
	//favBook = 12

	var otherfavBook string = "Digital Minimalism"
	fmt.Println(otherfavBook)

	var thirdFavBook string
	thirdFavBook = "I am Malala"
	fmt.Println(thirdFavBook)

	var myName string
	var myAge int
	var amICool bool
	fmt.Println(myName, myAge, amICool)
	fmt.Println("Length of myName =", len(myName))

	// var favNumber = 42
	// var favChocolate = "KitKat"
	// var favTeam = "CSK"

	// Compound creation (creating multiple variables) //
	// var favNumber, favChocolate, favTeam = 42, "KitKat", "CSK"

	// Block creation //
	var (
		favNumber    = 42
		favChocolate = "KitKat"
		favTeam      = "CSK"
	)
	fmt.Println(favNumber, favChocolate, favTeam)

	//var favAnimal = "Dog"

	// Variable declaration and assignment in one go //
	favAnimal := "Tiger"
	fmt.Println(favAnimal)

	pet1, pet2, pet3 := "Dog", "Cat", "Rat"
	fmt.Println(pet1, pet2, pet3)

	//pet3 := "test" // This will not work as there are no new variables
	pet3, pet4 := "Parrot", "Fish" // This will work as pet4 is a new variable, pet3 will be reassigned as well
	fmt.Println(pet3, pet4)

	// Constants //
	const myName2 = "Santh"
	// myName2 = "Vel" // This will throw an error as myName2 is defined as a constant
	fmt.Println(myName2)
}
