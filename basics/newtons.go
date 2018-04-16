package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
  z := x / 2
  for i := 1; i <= 10; i++ {
//	fmt.Println(z, " -= (", z, " * ", z, " - ", x, ") / ( 2 * ", x, ")")
	z -= (z*z - x) / (2*x)
//	fmt.Println("Iteration", i, " z val:", z)
  }
  return z
}

func main() {
	fmt.Println("Local:", Sqrt(7), "Math: ", math.Sqrt(7))
}
