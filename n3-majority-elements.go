package main

import "fmt"

func main_n3_majority_elements() {
	nums := []int{2, 1, 1, 3, 1, 4, 5, 6}
	result := []int{}

	thresold := len(nums) / 3
	c1, c2 := 0, 0
	f1, f2 := 0, 0

	for _, e := range nums {
		if c1 == e {
			f1++
		} else if c2 == e {
			f2++
		} else if f1 == 0 {
			c1 = e
			f1 = 1
		} else if f2 == 0 {
			c2 = e
			f2 = 1
		} else {
			f1--
			f2--
		}

	}
	fmt.Println(c1, c2, f1, f2, thresold)

	f1, f2 = 0, 0

	for _, e := range nums {
		if c1 == e {
			f1++
		} else if c2 == e {
			f2++
		}

	}
	if f1 > thresold {
		result = append(result, c1)
	}
	if f2 > thresold {
		result = append(result, c2)
	}

	fmt.Println(result)
}
