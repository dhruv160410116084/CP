// https://leetcode.com/problems/set-matrix-zeroes/
package main

import (
	"fmt"
)

type void struct{}

var member void

// bruteforce approach
func main1() {
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

// optimized approach
func main() {
	// matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	fmt.Println("Current Matrix: ")
	fmt.Println(matrix)

	rows := len(matrix)
	cols := len(matrix[0])
	col0 := 1

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}
	fmt.Println("Updated Matrix: ")
	fmt.Println(matrix)

	for i := rows - 1; i > -1; i-- {
		for j := cols - 1; j > 0; j-- {
			fmt.Println(i, j, matrix[i][j])
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}

		}
		if col0 == 0 {
			matrix[i][0] = 0
		}
	}
	fmt.Println("Ans Matrix: ")

	fmt.Println(matrix)

}
