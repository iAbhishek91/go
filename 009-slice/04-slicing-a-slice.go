/*
Demonstration of slicing a slice
* slicing an slice
* determine the length and slice
* how to recreate the slice from the existing slice (removing the reference to the original array)
*/

package main


import "fmt"

func main() {
	// SCENARIO-1: slice a slice and understand new slice
	// step-1: creating a new slice
	slice1 := []int{1,2,3,4,5,6,7,8,9,0}
  printSlice(slice1)

	// step-2: create a new slice from slice1 starting from 3rd element till 4th element
	// slice[i, j]
	// i represent the starting index of the new slice from the new slice. hence slice1[i] == slice2[0]
	// j length from the original array. here 4 is the total length from slice1, but the slice starts from 3rd
	// hence length is 2
	slice2 := slice1[2:4] // value: 3,4
  printSlice(slice2)

	// SCENARIO-2: slice a slice with third element (capacity)
	//step-1: creating a new slice
	slice3 := []int{11,12,13,14,15,16,17,18,19,10}
	printSlice(slice3)

	// step-2: create a new slice with capacity from an existing slice.
	// we can see in SCENARIO-1, capacity of slice2 is automatically
	// determined the the compiler based on the original slice and user do not have any choice.
	// "slice4" is a new slice with starts from 4th element of existing array, ends at 8th element, with capacity 5.
	// slice[i,j,k]
	// i is the index to start the new slice.
	// j is the length from the original array
	// k is the capacity of the original array
	slice4 := slice3[3:7:8]
	printSlice(slice4)
}


func printSlice(slice []int) {
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("Length of slice: %d and Capacity of slice: %d\n", len(slice), cap(slice))
	for i, d := range slice {
		fmt.Printf("index: %d, data: %d\n", i, d)
	}
}