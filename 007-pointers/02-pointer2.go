/* Read the comment line for further understand of pointers */

package main

import "fmt"

func main() {
	a := 20
	fmt.Println(&a) // output: 0xc0000b4008

	a = 40
	fmt.Println(&a) // output: 0xc0000b4008

	addOne(&a)
	fmt.Println(a, &a) // output: 41 0xc0000b4008
}

func addOne(a *int) {
	*a++ // a++ will throw ERROR, 
	// because "a" is a pointer and pointer variable contains memory address and not the actual value.
	// hence, "*a" points to the value of the "a".

	fmt.Println(a, &a) // demonstration of above explanation
	// output: 0xc0000b4008 0xc0000ae020
}