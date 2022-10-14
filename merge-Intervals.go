package main

import (
	"fmt"
	"math"
	"sort"
)

// type Matrix [][]int

// func (m Matrix) Len() int { return len(m) }
// func (m Matrix) Less(i, j int) bool {
// 	return m[i][0] < m[j][0]
// }

// func (m *Matrix) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func main_merge_intervals() {
	matrix := [][]int{
		// {1, 3}, {2, 6}, {8, 10}, {15, 18},'
		// {1, 4}, {4, 5},
		// {1, 5},
		{1, 4}, {2, 3},
	}
	var ans [][]int
	var temp []int
	m := matrix
	// m := Matrix(matrix)
	// sort.Sort(&m)
	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
	// fmt.Println(m)
	temp = m[0]

	for i := 1; i < len(m); i++ {
		if temp[1] >= m[i][0] {
			temp[1] = int(math.Max(float64(m[i][1]), float64(temp[1])))
		} else {
			ans = append(ans, temp)
			temp = m[i]
		}
	}
	ans = append(ans, temp)

	// for i, d := range m {
	// 	fmt.Println(i, d)
	// 	temp = d
	// }
	fmt.Println(ans)
}
