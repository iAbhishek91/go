/* demonstration of user defined type */
/*
two way to have user defined type are:
1. struct
2. using existing primitive types
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
  // STRUCT
	// 1. define a struct
	type employee struct {
		name string
		designation string
		age uint8
	}

	// 2. declare a variable of type struct
	var abhishek employee
	fmt.Println(abhishek)

	// 3. declaring and initialize the variable at the same time
	rob := employee{
		name: "Rob",
		designation: "frontend developer",
		age: 23,
	}
	bob := employee{"Bob", "Backend Developer", 25}
	fmt.Println(rob, " | ", bob)

	// 4. Updateing the value in struct
	abhishek.name = "Abhishek Das"
	abhishek.designation = "DevOps Engineer"
	abhishek.age = 28
  fmt.Println(abhishek)


	// EXISTING PRIMITIVE TYPE
	// 1. create a new type from existing
	type age uint8
	
	// 2. define a variable of new type
	var abhishekAge age
	fmt.Println(abhishekAge)
	fmt.Println(reflect.TypeOf(abhishekAge)) //main.age
	fmt.Println(reflect.TypeOf(abhishekAge).Kind()) // uint8
	fmt.Println(reflect.ValueOf(abhishekAge).Kind()) // unit8
	
	// 3. assign a value
	abhishekAge = 28
	fmt.Println(abhishekAge)

	// 4. assinging a variable of base type
	var a = uint8(5)
	// abhishekAge = a // ERROR: as uint8 cant be assigned to age type
	abhishekAge = age(a) // type conversion
	fmt.Println(abhishekAge) 
}