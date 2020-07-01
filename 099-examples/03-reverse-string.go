/* Reverse a string */
// Complexity: O(n)
package main


import "fmt"


func main() {
	strToReverse := "Hi my name is Abhishek"
	revStr := reverseString(strToReverse)
	
	fmt.Println(revStr)
}

// writing or commenting the code is shows modernization
func reverseString(str string) string {
	// VALIDATE THEN INPUT

	// Not mutating the string
  length := len(str)
	runes := make([]rune, length)

	for i, rune := range str {
		length--
		runes[length] = rune
	}	

	return string(runes)
}
