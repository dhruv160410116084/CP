package main

import "fmt"

func main_daily_temperatures() {
	temperatures := []int{30, 60, 90}
	ans := make([]int, len(temperatures))
	stack := []int{}

	for i := 0; i < len(temperatures); i++ {

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[top] = i - top
		}
		stack = append(stack, i)

	}

	// for i := 0; i < len(stack); i++ {
	// 	ans = append(ans, 0)
	// }
	fmt.Println(ans)

}
