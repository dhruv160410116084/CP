// https://leetcode.com/problems/merge-sorted-array/
package main

import "fmt"

func main() {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	p1, p2, i := m-1, n-1, m+n-1

	for p2 >= 0 {
		if p1 >= 0 && nums2[p2] < nums1[p1] {

			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
		i--
	}

	// for i := 0; i < m; i++ {
	// 	if len(nums2) > 0 && nums1[i] > nums2[j] {
	// 		nums1[i], nums2[j] = nums2[j], nums1[i]
	// 		for k, t := j, j; k < n; k++ {
	// 			if nums2[t] > nums2[k] {
	// 				nums2[t], nums2[k] = nums2[k], nums2[t]
	// 				t++

	// 			}
	// 		}
	// 		// j++
	// 	}

	// }

	// nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1, nums2)
}

func first_approach() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	j := 0

	for i := 0; i < m; i++ {
		if len(nums2) > 0 && nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			for k, t := j, j; k < n; k++ {
				if nums2[t] > nums2[k] {
					nums2[t], nums2[k] = nums2[k], nums2[t]
					t++

				}
			}
			// j++
		}

	}

	nums1 = append(nums1[:m], nums2...)
	fmt.Println(nums1, nums2)
}
