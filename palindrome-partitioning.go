package main

import "fmt"

func isPalindrome(s string, start int, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string, ans []string, i int) {
	if i >= len(s) {
		fmt.Println(ans)
		return
	}

	for j := i; j < len(s); j++ {
		if isPalindrome(s, i, j) {
			fmt.Println("palindrome: ", i, j)
			ans = append(ans, s[i:j+1])
			partition(s, ans, j+1)
			ans = ans[:len(ans)-1]
		}
	}
}

// a=97 z =122 A=65 Z=90
func main_palindrome_partition() {
	s := "Zzabaa"
	ans := []string{}
	fmt.Println(float64(1 / 2))
	fmt.Println(s[0])
	partition(s, ans, 0)
	// r := isPalindrome(s, 0, len(s)-1)
	// fmt.Println(r)
}
