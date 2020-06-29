/*
Problem: a array is give we need to find that addition of two number leads to sum(given)

Input: [1,2,4,4], 8
Output: True

> array is sorted
> Solution: 1. brute force O(n^2), using pointer O(n)
*/

package main

import "fmt"

var array [4]int = [4]int{4,4,4,5}
var sum = 8

func main() {
	fmt.Println(isSumAvailableInList(array, sum))
}


func isSumAvailableInList(array [4]int, sum int) bool {
	low := 0
	high := len(array) - 1

	for high > low {
		result := array[low] + array[high]
		fmt.Println(result)
		if result == sum {
			return true
		} else if result > sum {
			high --
		} else {
			low ++
		}
	}

	return false
}