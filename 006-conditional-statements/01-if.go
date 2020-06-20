/* Demonstration of different type of if statements 
* simple if
* if with multiple condition
* if with simple statement of evaluation
* if else ladder
*/

package main

import "fmt"

func main() {
	a := 5
	if a == 5 {
		fmt.Println("this is simple if")
	}


	c1 := 10
	c2 := 20
	if c1 >= 10 && c2 <= 20 { // we can use OR ||
		fmt.Println("if with multiple condition")
	}


  if a := return5(); a <= 5 {
		fmt.Println("if with simple statement of evaluation")
	}


	ie := 4
	if ie  < 4 {
		fmt.Println("if else ladder: ie < 4")
	} else if ie == 4 {
		fmt.Println("if else ladder: ie = 4")
	} else {
		fmt.Println("if else ladder: ie > 4")
	}
}

func return5() (int) {
	return 5
}