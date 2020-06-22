/* Demonstration of map iteration */

package main


import "fmt"


func main() {
	myMap := map[int]string{1: "red", 2: "black"}
	for k,v := range myMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}
