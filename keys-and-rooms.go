package main

import "fmt"

func bfs(node int, adjList map[int][]int, visited []int) {
	visited[node] = 1

	for _, v := range adjList[node] {
		if visited[v] != 1 {
			bfs(v, adjList, visited)
		}
	}
}

func main_keys_and_rooms() {
	rooms := [][]int{{1}, {2}, {3}, {}}
	visited := make([]int, len(rooms))
	adjList := make(map[int][]int)

	for i, v := range rooms {
		if len(v) > 0 {
			adjList[i] = append(adjList[i], v...)
		}
	}
	bfs(0, adjList, visited)
	// for _,v := range visited {
	// 	if v == 0 {
	// 		return false
	// 	}
	// }
	// return true
	fmt.Println(visited)
}
