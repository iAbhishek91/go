/* Demonstrate an array|slice can be converted to string */

package main

import (
	"strings"
	"fmt"
)

func main() {
	arr := []string{"a","b","c","d"}
	str := strings.Join(arr, "")

	fmt.Println(str)
}