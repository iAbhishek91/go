/* 
Demonstration of passing array in func
NOTE: as everything is pass by value in Go, passing array is expensive in terms of resources.
*/

package main

import "fmt"

func main() {
	array := [...]int{0,1,2,3,4,5,6,7,8,9}

	for _, element := range array {
		fmt.Println(element) // 0,1,2,3,4,5,6,7,8,9
	}

	forEachAddOne(&array)

	for _, element := range array {
		fmt.Println(element) // 1,2,3,4,5,6,7,8,9,10
	}
}


func forEachAddOne(arr *[10]int) {
	for i, _ := range arr {
		arr[i]++
	}
	// Note: if you do the following it wont work
	// as element is a copy of the element value
	// range function do not return pointer

	for _, element := range arr {
		element ++ // no effect of the line
	}
}