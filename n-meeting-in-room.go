package main

import (
	"fmt"
	"sort"
)

type meet struct {
	start int
	end   int
}

func main_n_meting_in_room() {
	start := []int{1, 0, 3, 5, 8, 5}
	end := []int{2, 6, 4, 7, 9, 9}

	meetings := []meet{}
	acceptedMeetings := []meet{}

	for i, v := range start {
		t := meet{start: v, end: end[i]}
		meetings = append(meetings, t)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].end < meetings[j].end
	})

	// fmt.Println(meetings)

	acceptedMeetings = append(acceptedMeetings, meetings[0])

	for i := 1; i < len(meetings); i++ {
		if acceptedMeetings[len(acceptedMeetings)-1].end < meetings[i].start {
			acceptedMeetings = append(acceptedMeetings, meetings[i])
		}
	}
	//ans
	fmt.Println(acceptedMeetings)

}
