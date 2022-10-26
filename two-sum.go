package main

import "fmt"

func main_two_sum() {
	nums := []int{3, 2, 4}
	target := 6
	result := []int{}

	m := make(map[int]int)

	for i, e := range nums {
		elem, ok := m[target-e]
		m[e] = i

		if ok && i != elem {
			result = append(result, i)
			result = append(result, elem)
			break
		}

	}

	// for k, v := range m {
	// 	// fmt.Println(k, v)
	// 	elem, ok := m[target-k]
	// 	if ok && v != elem {
	// 		result = append(result, v)
	// 		result = append(result, elem)
	// 		break
	// 	}
	// }
	fmt.Println(result)
}
