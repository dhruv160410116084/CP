package main

import "fmt"

func main_extract_digits() {
	nums := []int{13, 25, 83, 77}
	ans := []int{}

	for _, v := range nums {
		if v < 10 {
			// fmt.Println(v)
			ans = append(ans, v)
		} else {
			temp := []int{}

			for v > 0 {
				// fmt.Println(v % 10)
				temp = append(temp, v%10)
				v = v / 10

			}

			for i := len(temp) - 1; i >= 0; i-- {
				ans = append(ans, temp[i])
			}
		}
	}
	fmt.Println(ans)
}
