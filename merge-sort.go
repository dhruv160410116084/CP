package main

import "fmt"

func main_merge_sort() {
	arr := []int{5, 3, 2, 1, 4}
	fmt.Println(int(3 / 2))
	merge_sort(0, len(arr)-1, arr)
	fmt.Println(arr)
}

func merge(start int, mid int, end int, arr []int) []int {

	n1 := mid - start + 1
	n2 := end - mid

	L := make([]int, n1)
	R := make([]int, n2)
	copy(L, arr[start:start+n1+1])
	copy(R, arr[mid+1:mid+n2+1])

	x := 0
	y := 0
	m := start

	for x < n1 && y < n2 && m <= end {
		if L[x] <= R[y] {
			arr[m] = L[x]
			x++
		} else {
			arr[m] = R[y]
			y++
		}
		m++
	}

	for x < n1 {
		arr[m] = L[x]
		m++
		x++
	}

	for y < n2 {
		arr[m] = R[y]
		m++
		y++
	}
	return arr
}

func merge_sort(start int, end int, arr []int) {
	if start < end {
		mid := int((start + end) / 2)
		merge_sort(start, mid, arr)
		merge_sort(mid+1, end, arr)
		merge(start, mid, end, arr)

	}

}
