// Say that a "clump" in a list is a series of 2 or more adjacent elements
// of the same value. Return the number of clumps in the given list.
package main

import (
	"fmt"
)

func countClumps(s []int) int {
	// probably a way to not use both current and track ... think about it
	var clumps int
	var current int
	var track bool
	for i := 0; i < len(s); i++ {
		if track == false && s[i] == current {
			track = true
			clumps += 1
		} else if s[i] != current {
			track = false
			current = s[i]
		}
	}

	return clumps
}

func main() {
	fmt.Println(countClumps([]int{1, 2, 2, 3, 4, 4}))
	fmt.Println(countClumps([]int{1, 1, 2, 1, 1}))
	fmt.Println(countClumps([]int{1, 1, 1, 1, 1}))
	fmt.Println(countClumps([]int{2, 3, 3, 5, 6, 6, 7, 6, 6, 1, 8}))
}
