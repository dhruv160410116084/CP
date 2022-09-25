// https://leetcode.com/problems/set-matrix-zeroes/
package main

import (
	"fmt"
)

type void struct{}

var member void

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// fmt.Println(matrix, len(matrix), len(matrix[0]))
	rows := len(matrix)
	cols := len(matrix[0])
	row_set := make(map[int]void)
	col_set := make(map[int]void)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// fmt.Print(matrix[i][j])
			// fmt.Print(" ")
			if matrix[i][j] == 0 {
				// fmt.Println("zero detected at", i, "row", j, "col")
				// fmt.Println(reflect.TypeOf(strconv.Itoa(i)).String())
				// row_set[strconv.Itoa(i)] = member
				// col_set[strconv.Itoa(j)] = member
				row_set[i] = member
				col_set[j] = member
			}
		}
		// fmt.Println()
	}
	// fmt.Println(row_set, col_set)
	// fmt.Print(row_set[0])
	// _, ok := row_set[0]
	// fmt.Println(ok)

	for i := 0; i < rows; i++ {
		_, is_in_row := row_set[i]
		for j := 0; j < cols; j++ {
			_, is_in_col := col_set[j]
			if is_in_row || is_in_col {
				matrix[i][j] = 0
			}
		}
	}
	fmt.Println(matrix)
}
