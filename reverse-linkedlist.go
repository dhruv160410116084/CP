package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len  int
}

func (list *LinkedList) insert(data int) {
	if list.head == nil {
		list.head = &ListNode{data, nil}
		list.len++
	} else {
		itr := list.head
		for ; itr.Next != nil; itr = itr.Next {
		}
		itr.Next = &ListNode{data, nil}
		list.len++
	}
}

func (list *LinkedList) list() []int {
	nums := []int{}
	next := list.head

	for ; next != nil; next = next.Next {
		nums = append(nums, next.Val)
	}
	return nums
}

func (list *LinkedList) reverse() {
	var newhead *ListNode = nil

	for list.head != nil {
		next := list.head.Next
		list.head.Next = newhead
		newhead = list.head
		list.head = next
	}
	list.head = newhead
}

func main_reverse_linked_list() {
	var list LinkedList

	fmt.Println(list)

	// list.insert(4)

	nums := []int{3, 6, 8, 10}
	for _, e := range nums {
		list.insert(e)
	}
	list.reverse()
	numsl := list.list()
	fmt.Println(numsl)

}
