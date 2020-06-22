/* Demonstration of addition and deletion of map element */

package main

import "fmt"

func main() {
	myMap := map[int]string{1: "apple", 2: "mango"}

	// addition of a key
	myMap[3] = "grapes"

	for k,v := range myMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}

	//deletion of a key
	delete(myMap, 1)

	for k,v := range myMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}