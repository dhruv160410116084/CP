package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) insert(d int) {
	// fmt.Println("in head")
	if head.Val == -999 {
		// fmt.Println("in nil")
		*head = ListNode{d, nil}
		return
	}
	// fmt.Println(head)

	itr := head

	for itr.Next != nil {
		// fmt.Println(itr.Val, " ")
		itr = itr.Next
	}
	itr.Next = &ListNode{d, nil}

}

func (head *ListNode) iterate() {
	if head == nil {
		fmt.Print("[]")
		return
	}
	itr := head

	for itr != nil {
		fmt.Print(itr.Val, " ")
		itr = itr.Next
	}
	fmt.Println()
}

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
func (h1 *ListNode) add(h2 *ListNode) *ListNode {
	c := 0
	h := &ListNode{-9, nil}

	var head *ListNode
	// *head = *h
	for h1 != nil && h2 != nil {
		if h.Val == -9 {
			h = &ListNode{(h1.Val + h2.Val + c) % 10, nil}
			head = h
		} else {
			h.Next = &ListNode{(h1.Val + h2.Val + c) % 10, nil}
			h = h.Next
		}
		c = (h1.Val + h2.Val + c) / 10

		h1 = h1.Next
		h2 = h2.Next
	}

	for h1 != nil {
		if h.Val == -9 {
			h = &ListNode{(h1.Val + c) % 10, nil}
			head = h
		} else {
			h.Next = &ListNode{(h1.Val + c) % 10, nil}
			h = h.Next
		}
		c = (h1.Val + c) / 10
		h1 = h1.Next
	}

	for h2 != nil {
		if h.Val == -9 {
			h = &ListNode{(h2.Val + c) % 10, nil}
			head = h
		} else {
			h.Next = &ListNode{(h2.Val + c) % 10, nil}
			h = h.Next
		}
		c = (h2.Val + c) / 10
		h2 = h2.Next
	}
	if c != 0 {
		h.Next = &ListNode{c % 10, nil}
		h = h.Next
	}
	return head
}

func main_add_two_number_linkedList() {
	var head1 *ListNode = &ListNode{-999, nil}
	var head2 *ListNode = &ListNode{-999, nil}
	// var head3 *ListNode = &ListNode{-999, nil}

	head1.insert(3)
	head1.insert(7)
	// head1.insert(9)
	// head1.insert(9)
	// head1.insert(9)
	// head1.insert(9)
	// head1.insert(9)

	head2.insert(9)
	head2.insert(2)
	// head2.insert(9)
	// head2.insert(9)

	head1.iterate()
	head2.iterate()

	head3 := head1.add(head2)

	head3.iterate()

}
