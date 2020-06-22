/* Demonstration of validating a key exists */

package main

import "fmt"

func main() {
	// 1. validating existence of key using exists attribute
	myMap := map[string]int{"red": 1, "blue": 2}

	valueOfRed, redExists := myMap["red"]
	valueOfYellow, yellowExists := myMap["yellow"]
	fmt.Println(valueOfRed, redExists)
	fmt.Println(valueOfYellow, yellowExists)

	if redExists {
		fmt.Printf("value of red: %d\n", valueOfRed)
	}

	if yellowExists {
		fmt.Printf("value of yellow: %d\n", valueOfYellow)
	}

	// 2. validating existence of key using value. value is type zero value if key do not exists.
	value, _ := myMap["yellow"]
	if value != 0 {
		fmt.Printf("yellow exists")
	}

}
