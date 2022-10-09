package main

func swap(i int, j int, arr []int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
