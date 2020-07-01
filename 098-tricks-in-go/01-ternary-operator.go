/* There is NO ternary operator in Go */
/* Go suggest to use if else as a idiomatic way */
/* This is a short trick that is common in Go family */

package main

import "fmt"


func main() {
	a, b := 10, 20
	greater := (map[bool]int{true: a, false: b})[a > b]

	fmt.Println(greater)
}