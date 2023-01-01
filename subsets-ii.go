package main

import (
	"fmt"
	"sort"
)

func recursion_subset(nums []int, i int, ans []int, pre bool) {
	if i >= len(nums) {
		fmt.Println(ans)
		return
	}

	//not take
	recursion_subset(nums, i+1, ans, false)

	//take
	if i > 0 && nums[i] == nums[i-1] && !pre {
		return
	}
	ans = append(ans, nums[i])

	recursion_subset(nums, i+1, ans, true)

}

func main_subsets_ii() {
	nums := []int{2, 1, 2, 1, 3}
	ans := []int{}

	sort.Ints(nums)
	fmt.Println(nums)

	recursion_subset(nums, 0, ans, false)

}
