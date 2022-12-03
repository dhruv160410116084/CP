package main

import "fmt"

const n = 3

func recursion(arr []int, i int, set []int, ans *[][]int) {
	if i >= len(arr) {
		// fmt.Println(set)
		cpy := make([]int, len(set))
		copy(cpy, set)
		// fmt.Println(arr)
		*ans = append(*ans, cpy)
		// fmt.Println(ans)
		return
	}

	// choose
	set = append(set, arr[i])
	recursion(arr, i+1, set, ans)

	//not choose
	set = set[:len(set)-1]
	recursion(arr, i+1, set, ans)

	// return ans
}

func main_subset_sum() {
	arr := []int{1, 2, 2}
	set := []int{}
	ansarr := [][]int{}
	// fmt.Println(arr)

	recursion(arr, 0, set, &ansarr)
	fmt.Println(ansarr)
}
