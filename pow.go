package main

import "fmt"

func main_pow() {
	x := 2.00000
	n := -2
	nn := n
	ans := 1.0
	if nn < 0 {
		nn = -1 * nn
	}

	for nn > 0 {
		if nn%2 == 0 {
			x = x * x
			nn = nn / 2
		} else {
			ans = ans * x
			nn = nn - 1
		}
	}

	if n < 0 {
		ans = 1 / ans
	}

	fmt.Println(ans)

}
