/* Demonstration of addition and deletion of element from slice */
package main

import "fmt"

func main() {
	// SCENARIO-1: simple addition
	// Step-1: create a new slice, with first and last element initialized
	slice1 := []int{1, 4: 5}
	printSlice(slice1)
	// Step-2: initializing 2nd index
	slice1[1] = 2
	printSlice(slice1)
  // Step-3: this will append slice1
	slice1 = append(slice1, 6) // this will
	printSlice(slice1)


	//SCENARIO-2: validate the impact of append on other slice which refer to the same array
	// Step-1: create a new slice
	slice2 := []int{11,21,31}
	printSlice(slice2)

	// Step-2: slice slice1
	slice3 := slice2[1:3]
	printSlice(slice3)

	// Step-3: modify second value of slice2
	slice2[1] = 211

	// Step-4: validate change in slice3
	// the change is applied in slice3 as well because slice2 and slice3 reference to same array.
	printSlice(slice3)

	// Step-5: append slice3
	// A new array will be created as there is no capacity in slice3
	slice3 = append(slice3, 41)

	// Step-6 modify the third element of slice2
	// This change will not impact slice3 as before, as slice3 is in different array
	slice2[2] = 311
	printSlice(slice3)

	// Append two slice
	slice4 := []int{1,2}
	slice5 := []int{3,4}
	slice6 := append(slice4, slice5...)
	printSlice(slice6)
}


func printSlice(slice []int) {
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("Length of slice: %d and Capacity of slice: %d\n", len(slice), cap(slice))
	for i, d := range slice {
		fmt.Printf("index: %d, data: %d\n", i, d)
	}
}