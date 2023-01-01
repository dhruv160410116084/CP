package main

import "fmt"

func dfs(source int, destination int, visited []int, adjList map[int][]int) bool {
	visited[source] = 1
	res := false
	for _, v := range adjList[source] {
		if visited[v] != 1 {
			if v == destination {
				return true
			}
			res = dfs(v, destination, visited, adjList)
			if res {
				return true
			}
		}

	}
	return res
}

func main_find_if_path_exists_in_graph() {
	n := 10
	edges := [][2]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}
	source := 7
	destination := 5
	adjList := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}

	visited := make([]int, n+1)
	res := dfs(source, destination, visited, adjList)

	fmt.Println(res)
}
