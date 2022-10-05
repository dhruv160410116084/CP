// https://leetcode.com/problems/permutations/submissions/
package main

import (
	"fmt"
)

func listAllPermutation() {
	arr := []int{1, 2, 3}
	ans := [][]int{}
	fmt.Println(arr)
	generatePermutation(0, arr, &ans)
	fmt.Println(ans)

}

func swap(i int, j int, arr []int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func generatePermutation(start int, arr []int, ans *[][]int) {

	if start == len(arr)-1 {
		_arr := make([]int, len(arr))
		_cNum := copy(_arr, arr)
		fmt.Println("copy: ", _cNum)
		*ans = append(*ans, _arr)
	} else {
		for i := start; i < len(arr); i++ {
			swap(start, i, arr)
			generatePermutation(start+1, arr, ans)
			swap(start, i, arr)

		}
	}

}
