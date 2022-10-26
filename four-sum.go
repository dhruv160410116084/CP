package main

import (
	"sort"
)

func main_four_sum() {
	nums := []int{2, 2, 2, 2, 2}
	target := 8

	result := [][]int{}
	// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

	//a+b+c+d = target
	//c+D = target - (a+b)
	sort.Ints(nums)
	// fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			two_sum := nums[i] + nums[j]

			low, hi := j+1, len(nums)-1

			for low < hi {
				c_d := nums[low] + nums[hi]

				if c_d == (target - two_sum) {
					_temp := []int{nums[i], nums[j], nums[low], nums[hi]}
					result = append(result, _temp)
					for low < hi && nums[low] == _temp[2] {
						low++
					}
					for low < hi && nums[hi] == _temp[3] {
						hi--
					}
				} else if low < hi && c_d < (target-two_sum) {
					low++
				} else {
					hi--
				}

			}
			for j+1 < len(nums)-1 && nums[j] == nums[j+1] {
				j++
			}

		}
		for i+1 < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	// fmt.Println(result)

}
