/* demstrate how float literal are expressed */

package main

import "fmt"

func main() {
	f1 := 078.65465 // 78.65
	f2 := 1e3 // 1000
	f3 := 1_5 // 15
  f4 := 0.15e+0_2 // 15 

	fmt.Println(f1, f2, f3, f4)
}