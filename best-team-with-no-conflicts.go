package main

import (
	"fmt"
	"sort"
)

type Team struct {
	scores []int
	ages   []int
}

func (a Team) Len() int { return len(a.ages) }
func (a Team) Swap(i, j int) {
	a.ages[i], a.ages[j] = a.ages[j], a.ages[i]
	a.scores[i], a.scores[j] = a.scores[j], a.scores[i]
}
func (a Team) Less(i, j int) bool { return a.ages[i] < a.ages[j] }

func recurScore(i int, players Team, prev int) int {
	if i >= len(players.ages) {
		return 0
	}

	two := 0
	// fmt.Println("i: ", i, "one: ", one, "two: ", two)

	if prev == -1 || players.scores[prev] <= players.scores[i] {
		two = players.scores[i] + recurScore(i+1, players, i) //select
		temp := recurScore(i+1, players, prev)
		if temp > two {
			two = temp
		}
		return two
	}
	//  else if i == 0 {
	// 	fmt.Println("in else if ")
	// 	two = players.scores[i] + recurScore(i+1, players)

	// }
	return recurScore(i+1, players, prev)

}

func main_best_team() {

	scores := []int{4, 5, 6, 5}
	ages := []int{2, 1, 2, 1}

	players := Team{scores: scores, ages: ages}

	sort.Sort(Team(players))

	fmt.Println(players)

	result := recurScore(0, players, -1)
	fmt.Println("result:", result)

}
