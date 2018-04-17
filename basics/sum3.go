// given an array of integers, find three elements which add to a given sum.

package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 10
  found := false
	a := []int{-3, 2, 7, 1, 6, -5, 9, 21, -11}
	// step one: sort
	sort.Ints(a)

	// conversion of existing Python code.
	// research to see what is more Go idomatic?
	for i := 0; i < len(a)-2; i++ {
		l := i + 1
		r := len(a) - 1

    // breaks because I need to rewrite this as a function outside main()
    // otherwise this returns all triplet matches
    if found {
      break
    }

		for l < r {
			if a[i]+a[l]+a[r] == target {
				fmt.Printf("Triplet is %v, %v, %v\n", a[i], a[l], a[r])
        found = true
        break
			} else if a[i]+a[l]+a[r] < target {
				l += 1
			} else {
				r -= 1
			}
		}
	}
  if found == false {
    fmt.Println("None found.")
  }

}
