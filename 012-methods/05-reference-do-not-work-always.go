/* Demonstration of go reference do not work */
/* this program is similar to 04 */
/* Very important program to understand */

package main

import "fmt"

type age uint8

func (a *age) print() {
	// VERY VERY IMPORTNAT:
	// If a is struct it print the value
	// for primitive type, it would print the address of age.
	fmt.Println(a)

	// this will print the value of address in "a"
	fmt.Println(*a)
}


func main() {
	// IMPORTANT to understand
	// Below line is not same as separating it in two
	// age(20).print() // ERROR: cannot call pointer method
	// https://stackoverflow.com/questions/62589461/golang-auto-referencing-not-always-possible-to-get-the-address-of-a-value

	// however the below one works, but it prints the address and do not reference the value
	// It happens for all they type, pattern of display is different for struct.
	myAge := age(20)

	myAge.print()
	// Output:
	// 0xc0000b4002
	// 20
}