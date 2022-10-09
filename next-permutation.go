// https://leetcode.com/problems/next-permutation/
package main

import (
	"fmt"
	"sort"
)

func main_next() {
	nums := []int{3, 2, 1, 4, 5}
	// nums := []int{1, 3, 2}
	// sort.Sort(sort.Reverse(sort.IntSlice(nums[0:])))
	// fmt.Println(nums)

	bp1 := -1
	bp2 := -1

	for i := len(nums) - 2; i > -1; i-- {
		if nums[i] < nums[i+1] {
			bp1 = i
			break
		}
	}

	if bp1 != -1 {
		for i := len(nums) - 1; i > -1; i-- {
			if nums[bp1] < nums[i] {
				bp2 = i
				break
			}
		}
		// fmt.Println(bp2)

		temp := nums[bp1]
		nums[bp1] = nums[bp2]
		nums[bp2] = temp

		// fmt.Println(nums)
		// swap(bp1,bp2,nums)
	}

	// fmt.Println(bp1 + 1)
	if bp1 != -1 {
		// sort.Sort(sort.Reverse(sort.IntSlice(nums[bp1+1:])))
		sort.Ints(nums[bp1+1:])
	} else {
		sort.Sort(sort.IntSlice(nums))
	}

	fmt.Println(nums, bp1, bp2)

}
