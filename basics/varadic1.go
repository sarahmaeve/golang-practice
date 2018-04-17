package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	var total int

  // for number just adds up the index :(
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sum(5, 23, 94, 0, 2, -29, 38))
}
