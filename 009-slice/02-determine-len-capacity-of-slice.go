/* Determine length and capacity of slice */
// how length and capacity is calculated is mentioned in details of 04-slicing-a-slice.go

package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	fmt.Printf("Length: %d, and Capacity: %d\n", len(slice), cap(slice))
}
