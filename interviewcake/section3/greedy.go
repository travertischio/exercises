package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxProfit([]int{10, 7, 5, 8, 11, 9}))
}

func getMaxProfit(stockPrices []int) int {
	lowPrice := stockPrices[0]
	profit := stockPrices[1] - stockPrices[0]

	for i := 1; i < len(stockPrices); i++ {
		currentPrice := stockPrices[i]
		currentProfit := currentPrice - lowPrice

		if currentProfit > profit {
			profit = currentProfit
		}

		if currentPrice < lowPrice {
			lowPrice = currentPrice
		}
	}

	return profit
}
