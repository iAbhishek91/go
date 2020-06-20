/* demonstration of different type of variable declaration */

package main

import "fmt"

func main() {
	// shorthand variable declaration
	// allowed only inside a function
	a := 10
	//a := 20 // ERROR: re declaration is variable is not allowed in same block
	a = 20 // re assignment is allowed

	a1, a2 := 10, 5
	// a2 is possible to re declare
	// type should be same
	// one of the variable is new, in this case a3
	a3, a2 := 11, 80

	// var b // ERROR: this is not allowed, either value or type is required
	var b = "implicitly type is determined,"
	var c string = "explicitely mentioned,"
	var d, e float64
	// var d1 = nil // ERROR: not allowed to directly use nil of undefined type.
	// var f float32, g float64 // ERROR: not allowed to declare multiple type in one statement see below
	var (
		f float32
		g float64
	)
	// var b = "new value" // ERROR: re declaration of variable is not allowed in same block

	// d = 40 //ERROR: Undefined, need to declare a variable before initializing

	k, l := true, "other type,"

	fmt.Println(a, a1, a2, a3, b, c, d, e, f, g, k, l)
}
