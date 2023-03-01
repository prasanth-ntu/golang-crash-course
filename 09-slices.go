package main

import (
	"fmt"
	"reflect"
)

func main() {
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
}
