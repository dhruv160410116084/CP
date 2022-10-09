package main

import (
	"fmt"
	"math"
)

func main_buy_sell_stock() {
	prices := []int{7, 1, 5, 3, 6, 4}
	minPrice := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		minPrice = int(math.Min(float64(minPrice), float64(prices[i])))
		maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-minPrice)))
	}
	fmt.Println(maxProfit)

}
