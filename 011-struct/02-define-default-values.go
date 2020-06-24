/* demonstration of setting default values for each struct type */
/* This can be visualized as constructor in obj oriented programming lang */

package main

import "fmt"


// define a struct
type simple struct {
	a int
	b string
}

// set a method with pointer receiver
func (s *simple) setDefault() {
	s.a = 1
	s.b = "a"
}

// set a method with pointer receiver, kind of parameterized constructor
func (s *simple) initSimple(a int, b string) {
	s.a = a
	s.b = b
}


func main() {
	// declare a variable of type simple
	var simpleInstance simple
	fmt.Println(simpleInstance)

	// set the default values
	simpleInstance.setDefault();
	fmt.Println(simpleInstance)
	
	// kind of parameterized constructor
	// set the default values
	simpleInstance.initSimple(-1, "000");
  fmt.Println(simpleInstance)
}
