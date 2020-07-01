/* how to determine a variable is array or not */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [...]string{"cat", "mat", "bat"}
	fmt.Println(reflect.ValueOf(array).Kind())
}