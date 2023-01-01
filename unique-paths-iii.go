package main

import "fmt"

var result = 0
var directions = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func backtrack(grid [][]int, i int, j int, count int, nonObs int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == -1 {
		return
	}
	// if i == 0 && j == 0 {
	// 	fmt.Println("---------")
	// }
	// fmt.Println(i, j, count, nonObs)

	if grid[i][j] == 2 {
		if count == nonObs {
			result++
		}
		return
	}

	grid[i][j] = -1

	//backtract
	for _, v := range directions {
		backtrack(grid, i+v[0], j+v[1], count+1, nonObs)
	}

	grid[i][j] = 0

}

func main_unique_paths_iii() {
	// grid := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}
	grid := [][]int{{0, 1}, {2, 0}}

	start_x := -1
	start_y := -1
	nonObs := 0
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				start_x = i
				start_y = j
				nonObs++

			}
			if grid[i][j] == 0 {
				nonObs++
			}
		}
	}

	backtrack(grid, start_x, start_y, count, nonObs)
	fmt.Println(result)
}
