package main

import "fmt"

func main_lcs() {
	nums := []int{0, -1}

	result := 0
	set := make(map[int]int)

	for i, e := range nums {
		set[e] = i + 1
	}

	for k, _ := range set {
		if _, ok := set[k-1]; !ok {
			ll := 1
			// for _, next := set[k+1]; next && k < len(nums); k++ {
			// 	// _, has := set[k+1]

			// 	fmt.Println(set[k+1], next)
			// 	ll++
			// }
			for set[k+1] != 0 {
				ll++
				k++
			}
			if ll > result {
				result = ll
			}
		}
	}
	fmt.Println(result)

}
