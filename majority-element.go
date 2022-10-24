package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}

	cur_ele := nums[0]
	freq := 1

	for i := 1; i < len(nums); i++ {
		if freq == 0 {
			cur_ele = nums[i]
		}

		if nums[i] == cur_ele {
			freq++
		} else {
			freq--
		}
	}

	fmt.Println(cur_ele)
}
