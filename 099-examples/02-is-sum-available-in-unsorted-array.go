/*
Problem: a array is give we need to find that addition of two number leads to sum(given)

Input: [1,2,4,4], 8
Output: True

> array is sorted
> Solution: 1. brute force O(n^2), using hash tables O(n)
*/

package main

import "fmt"

var array = [5]int{4,1,3,4,7}

var sum = 8

func main() {
	fmt.Println(isSumAvailableInList(array, sum))
}

func isSumAvailableInList(array [5]int, sum int) bool {
	// save the complements of element
	complement := make(map[int]bool)
	
	for _, data := range array {
		if complement[sum - data] == true {
			return true
		} else {
			complement[data] = true
		}
	}

	return false
}