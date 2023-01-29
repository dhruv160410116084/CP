package main

import (
	"fmt"
)

func main_dijkstra() {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1

	adjList := make(map[int][][2]int)
	dist := make([]int, n)
	for i := 0; i < len(dist); i++ {
		dist[i] = 99999
	}
	queue := [][3]int{}

	for _, v := range flights {
		adjList[v[0]] = append(adjList[v[0]], [2]int{v[1], v[2]})
	}

	queue = append(queue, [3]int{src, 0, 0}) // (node,dist,stops)
	dist[src] = 0

	for len(queue) != 0 {
		d := queue[0][1]
		n := queue[0][0]
		c_k := queue[0][2]

		queue = queue[1:] //remove

		for _, v := range adjList[n] {
			fmt.Println("before: ", v, dist, "d: ", d, "n: ", n, "k: ", c_k)
			if d+v[1] < dist[v[0]] && c_k <= k {
				dist[v[0]] = d + v[1]
				if v[0] != dst {
					queue = append(queue, [3]int{v[0], d + v[1], c_k + 1})

				} else {
					queue = append(queue, [3]int{v[0], d + v[1], c_k})
				}
				fmt.Println("after: ", v, dist)

			}
		}

	}

	fmt.Println(dist[dst])

}
