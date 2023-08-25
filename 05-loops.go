package main

import "fmt"

func main() {

	fmt.Println("For loop")
	for i := 1; i <= 10; i++ {
		fmt.Println("i=", i)
	}

	fmt.Println("\nFor loop without increment/decrement operator")
	j := 1
	for j <= 10 {
		fmt.Println("j=", j)
		j++
	}

	fmt.Println("\nFor loop without increment/decrement operator and without stopping condition")
	k := 1
	for {
		fmt.Println("k=", k)
		k++
		if k > 10 {
			break
		}
	}

	fmt.Println("\nFor loop that prints only odd numbers")
	r := 1
	for r <= 20 {
		// Skip if even number
		if r%2 == 0 {
			r++
			continue
		}
		fmt.Println("r=", r)
		r++
	}
}
