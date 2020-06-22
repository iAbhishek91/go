/* Demonstrating different way of declaring and initializing slice */

package main

import (
	"fmt"
)

func main() {
	// 1. using var to declare a slice of same length and capacity
	var slice1 = make([]int, 5)

	// 2. using var to declare a slice of length 3 and capacity 5
	var slice2 = make([]int, 3, 5)

	// 3. using var and new to declare slice of length 3 and capacity 5
	var slice3 = new([5]int)[2:5] //here we are deciding the base array to be used by slice.

	// 4. declare a nil array, where the slice do not point to any array.
	var slice4 []int

	// 5. declare and initialize slice using the var keyword
	var slice5 []int = []int{101,102,103}

	// 7. using slice literal to declare a empty slice (the array is of size zero)
	slice7 := []int{}

	// 8. literal to declare and initialize an array
	slice8 := []int{11,22,33}
	
	// 9. as we can see in literal we can't set the length and capacity, but see how we can do that
	// declaring slice of length and capacity of 10
	slice9 := []int{9: 101010}

	printSlice(slice1)
	printSlice(slice2)
	printSlice(slice3)
	printSlice(slice4)
	printSlice(slice5)
	printSlice(slice7)
	printSlice(slice8)
	printSlice(slice9)
}


func printSlice(slice []int) {
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("Length of slice: %d and Capacity of slice: %d\n", len(slice), cap(slice))
	for i, d := range slice {
		fmt.Printf("index: %d, data: %d\n", i, d)
	}
}