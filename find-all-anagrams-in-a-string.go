package main

import "fmt"

func main() {
	s := "abab"
	p := "ab"
	count := 0
	result := []int{}

	sf := make([]int, 26)
	pf := make([]int, 26)

	for i := 0; i < len(p); i++ {
		sf[s[i]-'a']++
		pf[p[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if sf[i] == pf[i] {
			count++
		}
	}

	for i := 0; i < len(s)-len(p); i++ {
		if count == 26 {
			result = append(result, i)
		}

		l := s[i] - 'a'
		r := s[len(p)+i] - 'a'

		sf[r]++
		if sf[r] == pf[r] {
			count++
		} else if sf[r] == pf[r]+1 {
			count--
		}

		sf[l]--
		if sf[l] == pf[l] {
			count++
		} else if sf[l] == pf[l]-1 {
			count--
		}
	}

	if count == 26 {
		result = append(result, len(s)-len(p))
	}
	fmt.Println(result)
}
