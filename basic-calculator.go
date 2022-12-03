package main

import (
	"container/list"
	"fmt"
	"unicode"
)

type pair struct {
	sum  int
	sign int
}

func main_basic_cal() {
	s := "(1+(4+5+2)-3)+(6+8)"
	sum := 0
	sign := +1

	stack := list.New()

	for i := 0; i < len(s); i++ {

		if unicode.IsDigit(rune(s[i])) {
			// fmt.Println(i, string(s[i]))
			num := 0
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				num = (num * 10) + int(s[i]) - '0'
				i++
			}
			i--
			sum += (num * sign)
			sign = +1
		} else if s[i] == '(' {
			stack.PushBack(pair{sum, sign})

			sum = 0
			sign = +1

			// fmt.Println(stack.Peek())
		} else if s[i] == ')' {
			p := stack.Back()
			fmt.Println(p.Value)
			// sum = p.(pair).sum + (sum * p.(pair).sign)
			sum = p.Value.(pair).sum + (sum * p.Value.(pair).sign)
			stack.Remove(p)
		} else if s[i] == '-' {
			sign = (-1 * sign)
		}
	}
	fmt.Println(sum)
}
