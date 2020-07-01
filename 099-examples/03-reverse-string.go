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

	var revStr string

	for i := 0; i < len(str) ; i++ {
		// this solution is not best one at least from performance activities
		// strings are IMMUTABLE https://golang.org/ref/spec#String_types
		revStr = revStr + string(str[len(str) - i - 1])

	}

	return revStr
}
