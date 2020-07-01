/* Demonstration of string function works with range */

package main

import "fmt"

func main() {
	str := "abhishek is learning Go!"
	
	length := len(str)

  runes := make([]rune, length)

	for i, rune := range str {
		fmt.Println(i, string(rune))

		length--
		runes[length] = rune
	}
}