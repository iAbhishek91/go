/* Two sorted array is merged and sorted again */

package main

import "fmt"
import "sort"

func main() {
	firstArr := []int{1,2,3,4}
	secondArr := []int{3,4,6}

  // since the data are of type slice we are not using address
	mergeAndSortUsingDefinedFunction(firstArr, secondArr)
}


func mergeAndSortUsingDefinedFunction(firstArr []int, secondArr []int) {
	// VALIDATE the input

	// MERGE both the slices
	mergedAndSortedArr := append(firstArr, secondArr...)
	sort.Ints(mergedAndSortedArr)
	
	
	// for i, element
	fmt.Println(mergedAndSortedArr)
}