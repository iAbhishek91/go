/* Demonstration of multi-dimension array using matrix */

package main

import "fmt"

func main() {
	matrix1 := [2][2]int{{1, 1}, {1, 1}}
	matrix2 := [2][2]int{{2, 2}, {2, 2}}

	matrix3 := addMatrix(&matrix1, &matrix2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matrix3[i][j])
		}
		fmt.Println("")
	}
}

func addMatrix(matrix1 *[2][2]int, matrix2 *[2][2]int) [2][2]int {
	result := [2][2]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return result
}
