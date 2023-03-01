package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("You are allowed to enter")
	} else {
		fmt.Println("You are not allowed to enter")
	}

	if age >= 18 {
		fmt.Println("You are allowed to enter")
	} else if age >= 16 {
		fmt.Println("You are allowed to enter if you are accompanied by an adult")
	} else {
		fmt.Println("You are not allowed to enter")
	}
}
