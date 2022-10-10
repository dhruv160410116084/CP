// https://leetcode.com/problems/rotate-image/
package main

import (
	"fmt"
)

func main_rotate_matrix() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	// transpose

	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		// sort.Sort(sort.IntSlice(matrix[i]))
		// sort.Sort(sort.Reverse(sort.IntSlice(matrix[i])))
		for j := 0; j < len(matrix[i])/2; j++ {
			l := len(matrix)
			matrix[i][j], matrix[i][l-1-j] = matrix[i][l-1-j], matrix[i][j]
		}

	}

	fmt.Println(matrix)
}

/*
https://leetcode.com/problems/rotate-image/discuss/2506228/Golang-oror-Easy-to-Understand-oror-Faster-than-100

do checkout this solution

*/
