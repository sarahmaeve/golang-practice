package main

// Imports, formatting, basic syntax, calling functions from packages

import (
	"fmt"
	"math"
	"math/rand"
)

// simple func and variable declarations
func add(x, y float64) float64 {
  return x + y
}

func main() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("A number from 0 to 99:", rand.Intn(100)) // always 81 without seed
  var add1 float64 = 10.32
  add2 := 32.1042
  fmt.Println("My addition is:", add(add1, add2))
}
