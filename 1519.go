package main

import (
	"fmt"
	"math"
)

var ans []int

func dfsi(i int, adjMat map[int][]int, visited []bool, labels string) [26]int {
	visited[i] = true
	fmt.Println(i)
	counter := [26]int{}
	counter[labels[i]-'a'] += 1
	// fmt.Println(counter)

	for _, v := range adjMat[i] {
		if !visited[v] {
			childCounter := dfsi(v, adjMat, visited, labels)
			for _, v := range childCounter {
				// fmt.Println("i: ", i, v, counter, childCounter)
				counter[v] += childCounter[v]
			}
		}
	}
	// fmt.Println(counter)
	ans[i] = counter[i]
	return counter
}

func main_1519() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	labels := "abaedcd"

	adjMat := make(map[int][]int)
	visited := make([]bool, n)
	ans = make([]int, n)
	fmt.Println("hell")
	for _, v := range edges {
		adjMat[v[0]] = append(adjMat[v[0]], v[1])
		adjMat[v[1]] = append(adjMat[v[1]], v[0])
	}
	dfsi(0, adjMat, visited, labels)
	fmt.Println(ans, "---")

	fmt.Println(math.Log10E)

	f64 := float64(7.)/3 - float64(4.)/3 - float64(1.)
	fmt.Println(f64)

}
