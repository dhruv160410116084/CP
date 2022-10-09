// https://leetcode.com/problems/maximum-subarray/
package main

import (
	"fmt"
	"math"
)

func main_maximum_subarray() {
	nums := []int{5, 4, -1, 7, 8}

	max := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		max = int(math.Max(float64(sum), float64(max)))
		if sum < 0 {
			sum = 0
		}
	}

	fmt.Println(max)
}
