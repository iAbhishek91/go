/* Demonstration how to find type of struct */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// define struct
  type name struct {
		firstName string
		lastName string
	}

	// define variable of struct name
	var abhishek name

	// verify the type
	fmt.Println(reflect.TypeOf(abhishek)) // main.name
	fmt.Println(reflect.TypeOf(abhishek).Kind()) // struct
	fmt.Println(reflect.ValueOf(abhishek).Kind()) //struct
}
