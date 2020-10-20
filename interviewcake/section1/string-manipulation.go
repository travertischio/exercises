package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	meetings := []Meeting{{1, 2}, {3, 6}, {2, 5}, {7, 9}}
	fmt.Println(mergeRanges(meetings))

	fmt.Println(stringReverse("cash"))

	fmt.Println(reverseWords("hello sir this is the captain"))

	myOrder := []int{3, 4, 6, 10, 11, 15}
	alicesOrder := []int{1, 5, 8, 12, 14, 19}
	fmt.Println(mergeSortedArrays(myOrder, alicesOrder))

	takeOutOrders := []int{1, 3, 5}
	dineInOrders := []int{2, 4, 6}
	servedOrders := []int{1, 2, 4, 6, 5, 3}
	fmt.Println(checkServiceOrder(takeOutOrders, dineInOrders, servedOrders))
}

type Meeting struct {
	startTime int
	endTime   int
}

func mergeRanges(meetings []Meeting) []Meeting {
	// first sort meetings by their start time
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i].startTime < meetings[j].startTime
	})

	for i := 1; i < len(meetings); i++ {
		a := meetings[i-1]
		b := meetings[i]

		if b.endTime <= a.endTime {
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		} else if b.startTime <= a.endTime {
			meetings[i-1] = Meeting{a.startTime, b.endTime}
			meetings = append(meetings[:i], meetings[i+1:]...)
			i--
		}
	}

	return meetings
}

func stringReverse(input string) string {
	inputBytes := []byte(input)
	for i := 1; i <= len(inputBytes)/2; i++ {
		stash := inputBytes[i-1]
		inputBytes[i-1] = inputBytes[len(inputBytes)-i]
		inputBytes[len(inputBytes)-i] = stash
	}

	return string(inputBytes)
}

func reverseWords(words string) string {
	reversedString := stringReverse(words)

	reversedStringArray := strings.Split(reversedString, " ")
	for i := 0; i < len(reversedStringArray); i++ {
		reversedStringArray[i] = stringReverse(reversedStringArray[i])
	}

	return strings.Join(reversedStringArray, " ")
}

func mergeSortedArrays(first []int, second []int) []int {
	var result []int

	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			result = append(result, first[i])
			i++
		} else {
			result = append(result, second[j])
			j++
		}
	}

	if i == len(first) {
		result = append(result, second[j:]...)
	} else {
		result = append(result, first[i:]...)
	}

	return result
}

func checkServiceOrder(takeOutOrders []int, dineInOrders []int, servedOrders []int) bool {
	i := 0
	j := 0

	for _, e := range servedOrders {
		if takeOutOrders[i] == e {
			if i < len(takeOutOrders)-1 {
				i++
			}
		} else if dineInOrders[j] == e {
			if j < len(dineInOrders)-1 {
				j++
			}
		} else {
			return false
		}
	}

	return true
}
