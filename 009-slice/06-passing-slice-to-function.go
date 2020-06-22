/*
Demostrate how to pass slice in a function
*/


package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}

	addOneToEachElement(slice1)

	// NOTE: slice is pass by value, but it looks like its pass by reference.
	for i,d := range slice1 {
		fmt.Printf("index: %d, data: %d\n", i, d)
	}
}

func addOneToEachElement(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i]++
	}
}