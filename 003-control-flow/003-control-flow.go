package main

import "fmt"

func main() {
	fmt.Println("first line")

	// loop
	for i := 0; i < 10; i++ {
		//conditional
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	//invoke function
	foo()

	fmt.Println("last line")
}

func foo() {
	fmt.Println("I am in foo")
}