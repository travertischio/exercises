package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	numbers, _ := reader.ReadString('\n')
	numbers = strings.TrimSpace(numbers)
	numArray := strings.Split(numbers, " ")
	strN, strD := numArray[0], strings.TrimSpace(numArray[1])
	numerator, _ := strconv.ParseInt(strN, 10, 32)
	denominator, _ := strconv.ParseInt(strD, 10, 32)
	divisor, num := 1, numerator
	for ; num < denominator; divisor++ {
		num += numerator
	}
	if num != denominator {
		divisor--
		remainder := denominator - (num - numerator)
		fmt.Println(divisor, " Remainder ", remainder)
	} else {
		fmt.Println(divisor)
	}
}
