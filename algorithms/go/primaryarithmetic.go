package main

import (
  "fmt"
)

func reverse(s string) string {
    chars := []rune(s)
    for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
        chars[i], chars[j] = chars[j], chars[i]
    }
    return string(chars)
}

func carry(twoNumbers []string) {
  firstNumber := []rune(reverse(twoNumbers[0]))
  secondNumber := []rune(reverse(twoNumbers[1]))

  carrys, previousCarry := 0, 0
  longerLength := len(firstNumber)
  if len(firstNumber) < len(secondNumber) {
    longerLength = len(secondNumber)
  }

  for i := 0; i < longerLength; i++ {
    first, second := 0, 0
    if i < len(firstNumber) {
      first = int(firstNumber[i]) - 48
    }
    if i < len(secondNumber) {
      second = int(secondNumber[i]) - 48
    }

    if first + second + previousCarry > 9 {
      previousCarry = 1
      carrys += 1
    } else {
      previousCarry = 0
    }
  }

  fmt.Printf("%d\n", carrys)
}

func main() {
  carry([]string{"123", "456"})
  carry([]string{"555", "555"})
  carry([]string{"123", "594"})
  carry([]string{"9248", "832"})
}
