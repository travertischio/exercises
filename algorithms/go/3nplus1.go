package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for _, value := range lines {
		input := strings.Split(value, " ")
		strI, strJ := input[0], strings.TrimSpace(input[1])
		i, _ := strconv.ParseInt(strI, 10, 32)
		j, _ := strconv.ParseInt(strJ, 10, 32)

		maxCount := 0
		for k := i; k <= j; k++ {
			count := 1
			for l := k; l > 1; count++ {
				if l%2 == 0 {
					l = l / 2
				} else {
					l = l*3 + 1
				}
			}

			if count > maxCount {
				maxCount = count
			}
		}
		fmt.Println(i, j, maxCount)
	}
}
