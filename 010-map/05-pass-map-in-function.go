/* Demonstration of passing map in function */

package main

import "fmt"

func main() {
	myMap := map[int]string{1:"apple", 2:"banana"}

	newKey, newValue := 3, "mango"

	// Similar to slice map are pass by value but looks like pass by reference.
	// this is because both map and slice reference data type, which uses other data structure under the hood.
	addAKeyInMap(myMap, newKey, newValue)

	for k,v := range myMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}


func addAKeyInMap(myMap map[int]string, key int, value string) {
	myMap[key] = value
}