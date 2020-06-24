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
	if s.a == 0 { // apply the change if the value corresponds to zero value
		s.a = 1
	}
	if s.b == "" { // apply the change if the value corresponds to zero value
		s.b = "a"
	}
}


func main() {
	// declare a variable of type simple
	var simpleInstance simple
	fmt.Println(simpleInstance)

	// set the default values
	simpleInstance.setDefault();
	fmt.Println(simpleInstance)
}
