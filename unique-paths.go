package main

import "fmt"

// const m = 23
// const n = 12

func main_unique_paths() {
	m := 23
	n := 12

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	// fmt.Println(dp)
	for i, _ := range dp {
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	// fmt.Println(dp)
	ans := recur(0, 0, m, n, dp)
	fmt.Println(ans)
}

func recur(i int, j int, m int, n int, dp [][]int) int {
	if i >= m || j >= n {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}

	if dp[i][j] != -1 {
		// fmt.Println(dp[i][j], i, j)
		return dp[i][j]

	}
	//right
	right := recur(i, j+1, m, n, dp)
	//down
	left := recur(i+1, j, m, n, dp)
	dp[i][j] = left + right
	return dp[i][j]

}
