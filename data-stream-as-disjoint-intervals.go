package main

import (
	"fmt"
	"sort"
)

type SummaryRanges struct {
	arr map[int]int
}

func Constructor_sr() SummaryRanges {
	return SummaryRanges{arr: make(map[int]int)}
}

func (this *SummaryRanges) AddNum(value int) {
	// this.arr = append(this.arr, value)
	this.arr[value] = 1
}

func (this *SummaryRanges) GetIntervals() [][]int {
	intervals := [][]int{}
	arr := []int{}
	for k, _ := range this.arr {
		arr = append(arr, k)
	}

	if len(arr) == 1 {
		intervals = append(intervals, []int{arr[0], arr[0]})
	} else if len(arr) > 1 {
		sort.Ints(arr)
		for i := 0; i < len(arr); i++ {
			j := i

			for j+1 < len(arr) && arr[j]+1 == arr[j+1] {
				j++
			}
			intervals = append(intervals, []int{arr[i], arr[j]})
			i = j
		}
	}

	return intervals
}

func main_data_stream_as_disjoint_intervals() {
	summary := []string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"}
	qr := [][]int{{}, {1}, {}, {3}, {}, {7}, {}, {2}, {}, {6}, {}}

	var data SummaryRanges

	for i, v := range summary {
		if v == "SummaryRanges" {
			data = Constructor_sr()
			fmt.Println(nil)
		} else if v == "addNum" {
			data.AddNum(qr[i][0])
			fmt.Println(nil)
		} else if v == "getIntervals" {
			t := data.GetIntervals()
			fmt.Println(t)
		}
	}
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
