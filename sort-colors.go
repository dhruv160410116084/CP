// https://leetcode.com/problems/sort-colors/
package main

import "fmt"

func main_sort_colors() {
	nums := []int{2, 0, 2, 1, 1, 0}
	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			temp := nums[low]
			nums[low] = nums[mid]
			nums[mid] = temp
			low++
			mid++
		case 1:
			mid++
		case 2:
			temp := nums[mid]
			nums[mid] = nums[high]
			nums[high] = temp
			high--
		}
	}

	fmt.Println(nums)

}
