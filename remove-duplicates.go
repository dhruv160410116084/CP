package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	i := 0

	// k := 0

	// if len(nums) == 1 {
	// 	return 1
	// }

	for j := i + 1; j < len(nums); j++ {
		if nums[i] != nums[j] {

			nums[i+1] = nums[j]
			i++
			// k++
		}
	}
	fmt.Println(nums, k, i+1)
}
