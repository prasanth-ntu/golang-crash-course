package main

import "fmt"

func main() {
	// For Loop //
	for i := 1; i <= 10; i++ {
		fmt.Println("i=", i)
	}

	// for loop without increment/decrement operator //
	j := 1
	for j <= 10 {
		fmt.Println("j=", j)
		j++
	}

	// for loop without increment/decrement operator and without stopping condition //
	k := 1
	for {
		fmt.Println("k=", k)
		k++
		if k > 10 {
			break
		}
	}

	r := 1
	for r <= 100 {
		// Skip if even number
		if r%2 == 0 {
			r++
			continue
		}
		fmt.Println("r=", r)
		r++
	}
}
