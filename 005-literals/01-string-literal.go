/* how strings are represented in go language */

package main

import "fmt"

func main() {
	// only TWO type of string are supported
	// 1. interpreted (in double quotes)
	// 2. raw (in back quotes)
	// NOTE: single quotes are not stiring in Go. Look at rune literal

	str1 := "Abhishek" // interpreted

	str2 := `My 
name is Abhishek
Nothing else` //raw

	str3 := "My\nname is Abhishek\nNothing elses"

	fmt.Println(str1, str2, str3)
}