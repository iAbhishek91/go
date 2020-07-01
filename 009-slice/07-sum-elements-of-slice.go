package main

import "fmt"

func main() {
	slice := []int{1,2,3,4}
	fmt.Println(Sum(slice...))
}

// this is special way of passing argument to function.
// this function can take any number of argument.
// ...T should be the last argument.
func Sum(nums ...int) int {
	res := 0
	for _, num := range nums {
		res = res + num
	}

	return res
}