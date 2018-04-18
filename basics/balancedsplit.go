// Given a non-empty list, return true if there is a place to split the list
// so that the sum of the numbers on one side is equal to the sum of the numbers
// on the other side.

package main

import (
	"fmt"
	"sort"
)

func sumSlice(s []int) int {
	var total int
	for _, val := range s {
		total += val
	}
	return total
}

func canBalance(s []int) bool {
	// sort slice, then
	sort.Ints(s)
	for x := range s {
		if sumSlice(s[0:x]) == sumSlice(s[x:]) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canBalance([]int{1, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{2, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{10, 10}))
	fmt.Println(canBalance([]int{1, 5, 7, 3, 4, 20}))
}
