package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findOptimalMovies(248, []int{60, 140, 169, 94, 108}))

	fmt.Println(checkPalindrome("racecar"))

	fmt.Println(createWordCloud("We came, we saw, we conquered... then we ate Bill's (Mille-Feuille) cake."))
}

func findOptimalMovies(flightLength int, movieLengths []int) bool {
	remainingLengths := make(map[int]bool)

	for _, movieLength := range movieLengths {
		remainingLength := flightLength - movieLength
		if _, ok := remainingLengths[remainingLength]; ok {
			return true
		}

		remainingLengths[movieLength] = true
	}

	return false
}

func checkPalindrome(input string) bool {
	chars := []rune(input)
	appearances := make(map[rune]bool)

	for _, char := range chars {
		if _, ok := appearances[char]; ok {
			delete(appearances, char)
		} else {
			appearances[char] = true
		}
	}

	if len(appearances) > 1 {
		return false
	}

	return true
}

func createWordCloud(input string) map[string]int {
	result := make(map[string]int)

	words := strings.Split(strings.ToLower(input), " ")
	for _, word := range words {
		word = strings.Trim(word, "!.,;:?()")
		if _, ok := result[word]; ok {
			result[word]++
		} else {
			result[word] = 1
		}
	}

	return result
}
