/* demstrate how function literal are expressed */

package main

import "fmt"

func main() {
	// function literal are anonymous function
	// very very similar to javascript
	// they can be invoked in two way, 1. assign it to variable; 2. IIFE
	// for more info about function syntax, refer: function

	func(a int, b int) {
		fmt.Println("Sum is", a+b)
	}(3, 9) // example of IIFE

	cal := func(a int, b int) int { return a * b }
	res := cal(40, 9)
	fmt.Println("Multiplication result is", res)
}
