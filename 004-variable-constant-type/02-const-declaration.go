package main

import "fmt"

// it's ok to have two main for example purpose
// cant build this package with multiple main
func main() {
	const a int = 80
	
	// const b int // ERROR: value is missing
	const (
		b = 90
		c = "A"
	)

	// iota helps to define the sequence of constants
	const (
		day1 = iota
		day2
		day3
		day4
		day5
		day6
		day7
	)

	const (
		c1 = iota + 1
		c2
		c3
		c4
	)

	const (
		s1 = string(iota + 65)
		s2
		s3
		s4
	)

	fmt.Println(a, b, c)
	fmt.Println(day1, day2, day3, day4, day5, day6, day7) // 0, 1, 2, 3, 4, 5, 6, 7
	fmt.Println(c1, c2, c3, c4) // 1, 2, 3, 4
	fmt.Println(s1, s2, s3, s4) // A, B, C, D
}
