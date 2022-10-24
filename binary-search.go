package main

import "fmt"

func main_binary_search() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	start := 0
	end := len(arr) - 1
	target := 7
	found := -1

	for start <= end {
		mid := start + int((end-start)/2)
		if arr[mid] == target {
			found = mid
			break
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Println(found)
}
