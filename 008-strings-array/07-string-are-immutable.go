/* Understand the meaning of strings are immutable means :

In go doc it explains STRING as

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.
Strings are immutable: once created, it is impossible to change the contents of a string.

UNDERSTAND the above

String are saved in variable. The variable has address to the string and the length. [NOTE: measuring length is O(1)].
Hence the variable address remain same(because variable are mutable in Go), but the address changes every time when we change a string.

NOT RECOMMENDED

Don't to string manipulation which causes string to change, as its expensive, everytime we do that underlying address get changed.
New memory allocation, initialization, garbage collection of the old memory are all the task that is done under the hood, when we mutate a string.
*/

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "abhishek"
	fmt.Println("variable value:", str, "| variable address:", &str, "| headers", (*reflect.StringHeader)(unsafe.Pointer(&str)))

	str = "Hi!," + str
	fmt.Println("variable value:", str, "| variable address:", &str, "| headers", (*reflect.StringHeader)(unsafe.Pointer(&str)))
}
