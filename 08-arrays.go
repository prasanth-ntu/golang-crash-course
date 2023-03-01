package main

import "fmt"

func main() {
	// Approach 1 //
	purchases := [5]float32{20.1, 21.1, 22.2, 21.3, 22.4}
	fmt.Println(purchases)
	fmt.Println(purchases[0], purchases[4])

	// Approach 2 //
	var sales [4]float32 // Go assigns default value to each element in array
	fmt.Println(sales)
	sales[0] = 10.1
	fmt.Println(sales)

	// Approach 3 - Not specifying the size of the array while := //
	sales2 := [...]float32{10.1, 11.1, 12}
	fmt.Println(sales2)

	fmt.Println("Iterating through 'purchases' array using for loop")
	for i := 0; i < len(purchases); i++ {
		fmt.Println(i, purchases[i])
	}
}
