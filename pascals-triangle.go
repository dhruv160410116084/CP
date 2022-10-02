// https://leetcode.com/problems/pascals-triangle/
package main

import (
	"fmt"
)

func mainpascal() {
	const rows = 5
	arr := [][]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				arr = append(arr, []int{1})
			} else {
				if j == i {
					arr[i] = append(arr[i], 1)
				} else {
					arr[i] = append(arr[i], arr[i-1][j]+arr[i-1][j-1])
				}
			}

		}
	}
	fmt.Println(arr)
}
