package main

import (
	"fmt"
	"math"
)

func main_minimum_average_difference() {
	nums := []int{2, 5, 3, 9, 5, 3}

	prefix_sum := []int{nums[0]}
	size := len(nums)

	for i := 1; i < len(nums); i++ {
		prefix_sum = append(prefix_sum, prefix_sum[i-1]+nums[i])
	}

	// fmt.Println(prefix_sum)
	min_avg := int(math.Abs(float64(prefix_sum[size-1] / len(prefix_sum))))
	min_idx := size - 1
	fmt.Println(min_avg)

	for i := 0; i < len(prefix_sum)-1; i++ {
		// fmt.Println(i)
		avg := (prefix_sum[i] / (i + 1)) - ((prefix_sum[size-1] - prefix_sum[i]) / (size - (i + 1)))
		avg = int(math.Abs(float64(avg)))
		fmt.Println(avg)
		if avg < min_avg {
			min_avg = avg
			min_idx = i
		} else if avg == min_avg && i < min_idx {
			min_idx = i
		}
	}

	// n_avg := int(math.Abs(float64(prefix_sum[size-1] / len(prefix_sum))))
	// // n_avg := prefix_sum[size-1] / len(prefix_sum)
	// fmt.Println(n_avg)
	// if n_avg < min_avg {
	// 	min_avg = n_avg
	// 	min_idx = size - 1
	// }
	// min_avg = int(math.Abs(float64(min_avg)))
	fmt.Println("ans:", min_idx)
	// return min_avg
}
