package main

import "fmt"

func main() {
	var num1 = 4
	var num2 = 3

	// Arithmetic Operators (+, -, *, /, %) //
	var sum = num1 + num2
	var difference = num1 - num2
	var quotient = float32(num1) / float32(num2)
	var product = num1 * num2

	var remainder = num1 % num2

	fmt.Println(sum, difference, quotient, product, remainder)

	// Relational Operators (<, <=, >, >=, ==, !=) //
	var result = num1 > num2
	fmt.Println(result)

	result = num1 != num2
	fmt.Println(result)

	// Logical Operators (&&, ||, !) //
	const name = "Santh"
	var age = 30
	var result1 = (name == "Santh") && (age == 30)
	var result2 = !(name == "Vel") || (age <= 25)
	fmt.Println(result1, result2)

}
