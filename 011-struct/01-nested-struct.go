/* Determine declaration analysis of nested struct */

package main


import "fmt"


func main() {
  // define a struct
	type address struct {
		streetNumber uint8
		city string
		country string
	}

	type employee struct {
		name string
		addresses address
	}

	// declare and initialize a variable of type employee
	abhishek := employee{
		name: "Abhishek Das",
		addresses: address{
			streetNumber: 12,
			city: "London",
      country: "UK",
		},
	}
	fmt.Println(abhishek)
}