/* Two sorted array is merged and sorted again */

package main

import "fmt"
import "sort"

func main() {
	firstArr := []int{1,2,3,4}
	secondArr := []int{1,1,1,1,1} // 0,4,2

	// own logic
	fmt.Println(mergedAndSortedArr(firstArr, secondArr))
	fmt.Println("-------------------------------------")
	
	// since the data are of type slice we are not using address
	mergeAndSortUsingDefinedFunction(firstArr, secondArr)
}


func mergedAndSortedArr(firstArr []int, secondArr []int) []int {
	// Check input
	if len(firstArr) == 0 {
		return secondArr
	}
	if len(secondArr) == 0 {
	  return firstArr
	}
	
	// sort the array : we are sorting prior to merging as we want to merge it using custom logic
	// and not use sorting algorithm post merging.
	sortArray(firstArr)
	sortArray(secondArr)
	fmt.Println(firstArr, secondArr)

	// merge the array
	var mergedArr []int
	var elemToBeComparedArr1, elemToBeComparedArr2 int
	isFirstArrayComplete, isSecondArrayComple := false, false
	lengthOfArray1, lengthOfArray2 := len(firstArr), len(secondArr)
	currentIndexOfArr1 := 0
	currentIndexOfArr2 := 0

	for !isFirstArrayComplete && !isFirstArrayComplete {
		if currentIndexOfArr1 == lengthOfArray1 {
			isFirstArrayComplete = true
			break
		}
		if currentIndexOfArr2 == lengthOfArray2 {
			isSecondArrayComple = true
			break
		}

		elemToBeComparedArr1 = firstArr[currentIndexOfArr1]
		elemToBeComparedArr2 = secondArr[currentIndexOfArr2]
		if elemToBeComparedArr1 == elemToBeComparedArr2 {
			mergedArr = append(mergedArr, elemToBeComparedArr1, elemToBeComparedArr2)
			currentIndexOfArr1++
			currentIndexOfArr2++
		} else if elemToBeComparedArr1 > elemToBeComparedArr2 {
			mergedArr = append(mergedArr, elemToBeComparedArr2)
			currentIndexOfArr2++
		} else {
			mergedArr = append(mergedArr, elemToBeComparedArr1)
			currentIndexOfArr1++
		}
	}
	
	// appending all element if one of the array completes
	if (!isFirstArrayComplete) {
		mergedArr = append(mergedArr, firstArr[currentIndexOfArr1:lengthOfArray1]...)
	}
	if (!isSecondArrayComple) {
		mergedArr = append(mergedArr, secondArr[currentIndexOfArr2:lengthOfArray2]...)
	}


  // return result array
	return mergedArr
}


func sortArray(array []int) {
	// NOTE: we will see how to sort in later phases, while studing about sorting
	// Basically this are slices and hence are passed by reference : refer slice section.
	// This is reason of NO Retturn
	sort.Ints(array)
}



func mergeAndSortUsingDefinedFunction(firstArr []int, secondArr []int) {
	// VALIDATE the input

	// MERGE both the slices
	mergedAndSortedArr := append(firstArr, secondArr...)
	sort.Ints(mergedAndSortedArr)
	
	
	// for i, element
	fmt.Println(mergedAndSortedArr)
}