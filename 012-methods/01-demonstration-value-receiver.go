/* Demonstration of declaration of method, and invoking of method using value receiver */

// IMP: This makes the struct working with Primitive nature
package main

import "fmt"

type simple struct {
	a int
	b string
}

func (s simple) print() {
	fmt.Println(s)
}

func (s simple) setValue(a int, b string) {
	s.a = a
	s.b = b
	fmt.Println(s)
}

func (s simple) setValueNReturn(a int, b string) simple {
	s.a = a
	s.b = b
	return s
}



func main() {
	s := simple{10, "London"}
	
	s.print()
	// Output: {10 London}

	s.setValue(20, "UK")
	// Output {20 UK}

	s.print()
	// Output: {10 London}
	// value will not change automatically as we are using value receiver.

	s = s.setValueNReturn(30, "UK") 

	s.print()
	// Output: {30 UK}
	// as we can user need to manually need to apply the change on the type.
}