package main

import (
	"sort"
)

func main_2225() {
	matches := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}
	answer := [][]int{{}, {}}

	map_match := make(map[int]int)

	for _, m := range matches {
		// fmt.Println(m)

		map_match[m[1]] = map_match[m[1]] + 1
		if _, ok := map_match[m[0]]; !ok {
			map_match[m[0]] = 0

		}

	}
	// fmt.Println(map_match)

	for k, v := range map_match {
		if v == 0 {
			answer[0] = append(answer[0], k)
		} else if v == 1 {
			answer[1] = append(answer[1], k)

		}
	}
	sort.Ints(answer[0])
	sort.Ints(answer[1])

	// fmt.Println(answer)
}
