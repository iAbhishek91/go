/* demonstration of how Go can reference a variable 
when address is passed instead of the value. */

package main

import "fmt"


type simple struct {
	a int
	b string
}

func (s *simple) print() {
	fmt.Println(s)
	fmt.Println(*s)
}

func (s *simple) setValue(a int, b string) {
	s.a = a
	s.b = b
}


func main() {
	// declare a variable
	s := simple{1, "a"}

	// Go compiler will perform auto reference to the value 
	s.print()
	// Output:
	// &{1 a}
	// {1 a}

	// Go compiler will perform auto reference to the value
	s.setValue(2, "b")
	s.print()
	// Output: 
	// &{2 b}
	// {2 b}
}