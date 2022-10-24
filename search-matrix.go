package main

import "fmt"

func main_search_matrix() {
	arr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 16
	found := -1

	start := 0
	row_count := len(arr)
	col_count := len(arr[0])
	end := row_count*col_count - 1

	for start <= end {
		mid := start + int((end-start)/2)
		var r int = mid / col_count
		var c int = mid % col_count

		if arr[r][c] == target {
			found = mid
			break
		} else if arr[r][c] < target {

			start = mid + 1
		} else {
			end = mid - 1

		}

	}

	fmt.Println(found)
	// for i, num := range arr {
	// 	for j, ele := range num {
	// 		fmt.Println(i, j, ele, int(i/len(arr)), int(j/len(num)))

	// 	}
	// }
}
