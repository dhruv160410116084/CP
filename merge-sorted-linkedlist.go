package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (head *Node) create(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	tempHead := head
	head.data = nums[0]
	for i := 1; i < len(nums); i++ {
		head.next = &Node{nums[i], nil}
		head = head.next
	}
	return tempHead

}

func (head *Node) list() []int {
	if head == nil {
		return []int{}
	}
	nums := []int{}
	for head != nil {
		nums = append(nums, head.data)
		head = head.next
	}
	return nums
}

func (h1 *Node) merge(h2 *Node) *Node {
	if h1 == nil && h2 == nil {
		return nil
	} else if h1 == nil {
		return h2
	} else if h2 == nil {
		return h1
	}

	if h1.data > h2.data {
		temp := h1
		h1 = h2
		h2 = temp
	}
	head := h1
	var temp *Node

	for h1 != nil {
		if h1.data <= h2.data {
			temp = h1
			h1 = h1.next
		} else {
			temp.next = h2
			_temp := h1
			h1 = h2
			h2 = _temp
		}
	}

	temp.next = h2
	return head
}

func main_merge_sorted_linkedlist() {
	nums1 := []int{}
	nums2 := []int{1, 2, 5, 8, 10}

	ll1 := &Node{}
	ll2 := &Node{}

	head1 := ll1.create(nums1)
	head2 := ll2.create(nums2)

	head := head1.merge(head2)
	list := head.list()
	// _nums1 := head1.list()
	// _nums2 := head2.list()
	fmt.Println(list)
}
