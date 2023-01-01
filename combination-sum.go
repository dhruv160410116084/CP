package main

import (
	"fmt"
	"sort"
)

func recursion_combi(candidates []int, i int, sum int, target int, ans []int, pre bool) {
	if i >= len(candidates) {
		if sum == target {
			fmt.Println("---------", ans)
		}
		return
	}
	// fmt.Println("candidate: ", candidates[i], " i: ", i, " sum: ", sum, " ans: ", ans)

	if sum == target {
		fmt.Println("---------", ans)

		return
	}

	//not take
	// fmt.Println("not take")
	recursion_combi(candidates, i+1, sum, target, ans, false)

	if sum+candidates[i] <= target {
		// fmt.Println("take")
		if i > 0 && candidates[i] == candidates[i-1] && !pre {
			return
		}
		ans = append(ans, candidates[i])
		recursion_combi(candidates, i+1, sum+candidates[i], target, ans, true)
	}
}

func main_combination_sum() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	ans := []int{}
	sort.Ints(candidates)
	recursion_combi(candidates, 0, 0, target, ans, false)

}
