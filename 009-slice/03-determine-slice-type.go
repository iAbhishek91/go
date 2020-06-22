/* Demonstrating different way of declaring and initializing slice */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. using var to declare a slice of same length and capacity
	var slice1 = make([]int, 5)

	// know the type of the variable
	fmt.Println(reflect.ValueOf(slice1).Kind()) // remember the kind word from k8s
}