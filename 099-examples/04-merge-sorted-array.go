/* Two sorted array is merged and sorted again */

package main

import "fmt"

func main() {
	firstArr := []int{1,2,3,4}
	secondArr := []int{3,4,6}

	// remember everything is pass by value in Go, use pointers 
	mergeAndSort(&firstArr, &secondArr)
}


func mergeAndSort(firstArr []int, secondArr []int) {
	// VALIDATE the input

	largeArr := len(firstArr) > len(secondArr) ? firstArr : secondArr

	var mergedList []int
	
	for i, element
}