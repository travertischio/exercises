package main

import (
	"fmt"
	"math"
)

func lights(x float64) {
  fmt.Printf("%t\n", math.Mod(math.Sqrt(x), 1) == 0)
}

func main() {
	lights(10)
	lights(81)
	lights(3)
	lights(6241)
	lights(8191)
}
