/* demstrate how integer literal are expressed */

package main

import "fmt"

func main() {
	i1 := 077 // 78.65
	i2 := 0O76 // octal: 
	i3 := 0X11 // hexadecimal
  i4 := 9_8_9_7 //9897

	fmt.Println(i1, i2, i3, i4)
}