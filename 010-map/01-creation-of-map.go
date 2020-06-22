/* Demonstration of declaration and initialization of map */

package main


import "fmt"


func main() {
	//dict1 := map([]string)string{} // ERROR: key cant be slice

	// 1. Declare map using make (1, 3, 4) all are empty map
	dict1 := make(map[string]int)

	// 2. Declare and initialize map
	dict2 := map[string]int{"Red": 1, "Orange": 2}

	// 3. Declaring slice as map value
	dict3 := map[string][]int{}

	// 4. Creating and empty map. same as above
  dict4 := map[string]int{}

	// Creation of nil map are not allowed in Go
	// var dict5 map[string]int // ERROR: cant create nil map

	// 6. using var to declare and initialize map
	var dict5 = map[string]int{"Red": 1, "Orange": 2}

	printMap(dict1)
	printMap(dict2)
	printMap(dict4)
	printMap(dict5)

	fmt.Println(dict3)
}


func printMap(myMap map[string]int) {
	fmt.Println(myMap)
}