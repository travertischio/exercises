package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Scan input
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Iterate through array of input
	for index := int64(0); index < int64(len(lines)); {
		people, _ := strconv.ParseInt(lines[index], 10, 32)
		if people != 0 {
			nextTrip := index + people + 1
			thisTrip := lines[index+1 : nextTrip]
			index = nextTrip
			sum := 0.0
			for _, value := range thisTrip {
				val, _ := strconv.ParseFloat(value, 64)
				sum += val
			}
			thisTripAvg := sum / float64(people)

			changeDown := 0.0
			changeUp := 0.0
			for _, value := range thisTrip {
				val, _ := strconv.ParseFloat(value, 64)
				if val-thisTripAvg > 0.01 {
					changeDown += val - thisTripAvg
				} else {
					changeUp += thisTripAvg - val
				}
			}
			smallerChange := changeDown
			if changeUp < changeDown {
				smallerChange = changeUp
			}
			_, remainingDecimals := math.Modf(smallerChange * 100.0)
			if remainingDecimals > 0.0 {
				smallerChange -= 0.005 // Always round down
			}
			fmt.Printf("$%.2f\n", smallerChange)
		} else {
			index++
		}
	}
}
