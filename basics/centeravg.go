// Day 3 exercise from http://www.golang-book.com/guides/bootcamp/week-1/day-3
// Implement a centeredAverage function that computes the average
// of a list of numbers, but removes the largest and smallest values.

package main

import (
	"fmt"
	"sort"
)

func centeredAverage(s []float64) float64 {
	sort.Float64s(s)
	s = s[1 : len(s)-1]
	return s[len(s)/2]
}

func main() {

	fmt.Println(centeredAverage([]float64{1, 2, 3, 4, 100}))
}
