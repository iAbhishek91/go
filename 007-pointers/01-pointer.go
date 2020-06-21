/* Demonstration of pointers and how to pass pointers in function */

package main

import "fmt"

func main() {
	a := 10

	fmt.Printf("value of a = %d\n", a)

	passByValue(a)

	fmt.Printf("value of a post pass by val= %d\n", a)

	passByRef(&a)

	fmt.Printf("value of a post pass by ref= %d\n", a)
}

func passByValue(a int) {
	a = 20
}


func passByRef(a *int) {
	*a = 20
}