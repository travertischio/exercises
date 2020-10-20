package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getMaxProfit([]int{10, 7, 5, 8, 11, 9}))

	fmt.Println(highestProduct([]int{-10, -10, 1, 3, 2}))
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

func highestProduct(arrayOfInts []int) int {
	lowestInt := arrayOfInts[0]
	secondLowestInt := arrayOfInts[0]

	highestInt := arrayOfInts[0]
	secondHighestInt := arrayOfInts[0]
	thirdHighestInt := arrayOfInts[0]

	for i := 1; i < len(arrayOfInts); i++ {
		currentInt := arrayOfInts[i]
		if currentInt < secondLowestInt {
			if currentInt < lowestInt {
				secondLowestInt = lowestInt
				lowestInt = currentInt
			}
		} else if currentInt > thirdHighestInt {
			if currentInt > secondHighestInt {
				if currentInt > highestInt {
					secondHighestInt = highestInt
					highestInt = currentInt

				} else {
					secondHighestInt = currentInt
				}
				thirdHighestInt = secondHighestInt
			} else {
				thirdHighestInt = currentInt
			}
		}
	}

	sort.Ints(arrayOfInts)
	if secondLowestInt < 0 {
		if thirdHighestInt*secondHighestInt < lowestInt*secondLowestInt {
			return lowestInt * secondLowestInt * highestInt
		}
	}

	return thirdHighestInt * secondHighestInt * highestInt
}
