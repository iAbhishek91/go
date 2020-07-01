/* demstrate how rune literal are expressed */
/*
A letter of an ancient Greek alphabate or sometime roman alphabet as well.
*/

package main

import "fmt"

func main() {
	// rune literal are enclosed in single quotes
	// they changes the content to unicode equivalent
	a := 'A' //65
	//a = 'AA' // ERROR: too many characters.
	n := '\n'
	someUnicode := '\000'
	// invalidUnicode := '\U00110000' // ERROR invalid unicode

	fmt.Println(a, n, someUnicode)
}