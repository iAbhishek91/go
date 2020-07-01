/* Two sorted array is merged and sorted again */

package main

import "fmt"

func main() {
	firstArr := []int{1,2,3,4}
	secondArr := []int{3,4,6}

	// remember everything is pass by value in Go, use pointers in case of array
	printBiggerArray(&firstArr, &secondArr)
}


func printBiggerArray(firstArr *[]int, secondArr *[]int) {
	// VALIDATE the input

	// finding biggest array
	largeArr := (map[bool]*[]int{true: firstArr, false: secondArr})[len(*firstArr) > len(*secondArr)]

	fmt.Println(*largeArr)
}