package main

import (
	"fmt"
	"strconv"
)

func main_reverse_polish_notation() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	stack := []int{}

	for _, v := range tokens {
		//number
		if i, err := strconv.Atoi(v); err == nil {
			fmt.Printf("%q looks like a number.\n", v)
			stack = append(stack, i)
		} else { //operator
			fmt.Printf("%q looks like an operator.\n", v)
			op2 := stack[len(stack)-1]
			op1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if v == "+" {
				// ans, _ := strconv.Atoi(op1) + strconv.Atoi(op2)
				stack = append(stack, op1+op2)
				// fmt.Println(stack, op1, op2)
			} else if v == "*" {
				fmt.Println(op1 * op2)
				stack = append(stack, op1*op2)

			} else if v == "-" {
				stack = append(stack, op1-op2)

			} else if v == "/" {
				// fmt.Println(math.Floor(float64(op1)/float64(op2)), op1, op2)

				stack = append(stack, op1/op2)

			}
		}

	}
	fmt.Println(stack)

}
