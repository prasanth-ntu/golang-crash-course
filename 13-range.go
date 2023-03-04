package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Range with array of integers")

	ids := []int{1, 12, 34, 343, 34, 4}
	fmt.Println(reflect.TypeOf(ids), len(ids))

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	fmt.Println("Range with map/dict")

	emails := map[string]string{
		"Santh": "santh@gmail.com",
		"Sean":  "sean@gmail.com",
	}
	for key, value := range emails {
		fmt.Println(key, value)
	}
	for key := range emails {
		fmt.Println(key)
	}
	for _, value := range emails {
		fmt.Println(value)
	}

}
