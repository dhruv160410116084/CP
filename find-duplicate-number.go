package main

import "fmt"

func main_find_duplicate_number() {
	nums := []int{3, 1, 3, 4, 2}

	tor, hare := nums[0], nums[0]

	for {
		tor = nums[tor]
		hare = nums[nums[hare]]
		if tor == hare {
			break
		}
	}

	tor = nums[0]
	for tor != hare {
		tor = nums[tor]
		hare = nums[hare]

	}

	fmt.Println(hare)

}
