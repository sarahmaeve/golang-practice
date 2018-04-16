package main

// Imports, formatting, basic syntax, calling functions from packages

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("A number from 0 to 99:", rand.Intn(100)) // always 81 without seed
}
