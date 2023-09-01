package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		Arrays vs Slices
		- Arrays are fixed in size, whereas slices are dynamic (so, more flexible)
		- Arrays to Slices and vice versa
			- To convert an array to slice, we can use [<start>:<end>] operator. Both <start> and <end> are optional
			- To convert a slice to array, we can use a for loop or "copy" function
		- append() function is used to add elements to the end of slice
	*/

	// Array approach limitation: The size of the array is fixed //
	purchases := [5]float32{20.1, 21.1, 22.2, 21.3, 22.4}
	fmt.Println("purchases:", purchases)
	fmt.Println("TypeOf(purchases):", reflect.TypeOf(purchases))
	// purchases[5] = 10.2 // Will throw an error >>> index 5 out of bounds [0:5]

	mySlice := purchases[:] // Get a slice of the array and make the array size dynamic
	fmt.Println("\nmySlice:", mySlice)
	fmt.Println("TypeOf(mySlice):", reflect.TypeOf(mySlice))

	mySlice = append(mySlice, 10.2, 10.3, 10.4)
	fmt.Println("mySlice:", mySlice)
	fmt.Println("TypeOf(mySlice):", reflect.TypeOf(mySlice))

	myOtherSlice := purchases[:3] // [0:3] is same as [:3]
	fmt.Println("\nmyOtherSlice:", myOtherSlice)
	fmt.Println("TypeOf(myOtherSlice):", reflect.TypeOf(myOtherSlice))

	purchases2 := []float32{20.1, 21.1, 22.2, 21.3, 22.4} // Makes the array size dynamic
	fmt.Println("\npurchases2:", purchases2)
	fmt.Println("TypeOf(purchases2):", reflect.TypeOf(purchases2))

	combine := append(mySlice, myOtherSlice...)
	fmt.Println(combine)

	var myarray, myarray2 [5]float32

	// Convert slice to array: Approach 1
	for i := 0; i < len(purchases2); i++ {
		myarray[i] = purchases2[i]
	}
	fmt.Println("\nmyarray:", myarray)
	fmt.Println("TypeOf(myarray):", reflect.TypeOf(myarray))

	// Convert slice to array: Approach 1
	copy(myarray2[:], purchases2)
	fmt.Println("\nmyarray2:", myarray2)
	fmt.Println("TypeOf(myarray2):", reflect.TypeOf(myarray2))
}
