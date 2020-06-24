/* Demonstration of declaration of method, and invoking of method using pointer receiver */

package main

import "fmt"


type simple struct {
	a uint8
	b string
}

// expecting a pointer
func (s *simple) print() {
	fmt.Println(s)
}

// expecting a pointer
func (s *simple) setValue(a uint8, b string) {
	s.a = a
	s.b = b
}


func main(){
	// declaring a pointer variable of type simple
	// "s" is pointer variable as its contains address
	s := &simple{10, "abhishek"}
	
	// we are passing "*s", because "s" is a pointer and "*s" is the value
	( *s).print()
	// Output: &{10 abhishek}
  // invoking method with pointer receiver

	// we are passing "*s", because "s" is a pointer and "*s" is the value
	( *s).setValue(20, "London")
	( *s).print()
	// Output: &{10 abhishek}
	// Note: unlike value-receiver, the value gor changed eventhough we have not returned
	// anything from the "setValue()"
	// This behavior can be referred as non-primitive
}