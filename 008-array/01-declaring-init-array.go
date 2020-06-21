/* Demonstration of different type array declaration and initializing */

package main

import "fmt"

func main() {
	// 1. blank array using var keyword
	var array1 [5]int
	for i, data := range array1 {
		fmt.Printf("index: %d, value: %d\n", i, data) // all the value are zero as they are type int.
	}

	
	// 2. partial declaration of array using var keyword
	var array2 [3]float32 = [3]float32 {1.34, 23.8} // partial assignment
	for i, data := range array2 {
		fmt.Printf("index: %d, data: %f\n", i, data)
	}


	// 3. initialization of array in separate lines
	var array3 [2]float64
	array3[0] = 3.14
	array3[1] = 99.99
	for i, d := range array3 {
		fmt.Printf("index: %d, data: %f\n", i, d)
	} 

	
	// 4. using array var to declare partial array by position
	var array4 [3]string = [3]string{1: "middle element"}
	displayArray(&array4)


	// 5. using array literal to declare nill array
	array5 := [3]string{}
	displayArray(&array5)


	// 6. using array literal to declare and assign
	array6 := [3]string{"UK", "USA", "UAE"}
	displayArray(&array6)


	// 7. compiler to determine the size of the array
	array7 := [...]string{"cat", "mat", "bat"}
	displayArray(&array7)
}



func displayArray(array *[3]string) {
  for i, d := range array {
		fmt.Printf("index: %d, data: %s\n", i, d)
	}
}