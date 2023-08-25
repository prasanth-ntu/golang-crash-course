package main

import "fmt"

func countDown(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

func addNumber(num1 int, num2 int) int {
	// return num1 + num2
	sum := num1 + num2
	return sum
}

func main() {
	fmt.Println("Intro to functions")

	fmt.Println("Countdown from 10")
	countDown(10)
	fmt.Println("Countdown from 10")
	countDown(5)

	theSum := addNumber(5, 10)
	fmt.Println("Sum of two numbers:", theSum)
}
