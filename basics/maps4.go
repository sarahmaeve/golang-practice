package main

import (
	"fmt"
  "sort"
)

func find_smallest_manual(sl []int) int {
	var smaller int = sl[0]  // don't use zero value
	for _, y := range sl {
		if y < smaller {
			smaller = y
		}
	}
	return smaller
}

func find_smallest(sl []int) int {
  sort.Ints(sl)
  return sl[0]
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println("Manual: ", find_smallest_manual(x))
  fmt.Println("Sort: ", find_smallest(x))
}
