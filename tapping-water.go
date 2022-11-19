package main

import (
	"fmt"
	"math"
)

func main_tapping_water() {

	height := []int{4, 2, 0, 3, 2, 5}
	left_max := []int{}
	right_max := make([]int, len(height))
	sum := 0

	left_max = append(left_max, height[0])
	right_max[len(right_max)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		_max := height[i]
		if left_max[i-1] > _max {
			_max = left_max[i-1]
		}
		left_max = append(left_max, _max)
	}
	fmt.Println(left_max)
	for j := len(height) - 2; j >= 0; j-- {
		_max := height[j]
		if right_max[j+1] > _max {
			_max = right_max[j+1]
		}
		right_max[j] = _max
	}

	// fmt.Println(right_max)

	// for i, j := 0, len(right_max)-1; i < j; i, j = i+1, j-1 {
	// 	right_max[i], right_max[j] = right_max[j], right_max[i]
	// }

	// fmt.Println(right_max)
	for i := 0; i < len(height); i++ {
		sum += int(math.Min(float64(right_max[i]), float64(left_max[i]))) - height[i]
	}
	// fmt.Println(sum)

}
