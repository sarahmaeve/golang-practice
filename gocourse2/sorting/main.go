// code assignment / challenge
// sort slices of strings and integers in different ways
// while digging through sort documentation
// https://godoc.org/sort

package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	// sort three types of slices in order and reverse order
	studyGroup := people{"Xavier", "Logan", "Jean", "Ororo", "Kitty", "Illyana"}
	sl := []string{"Professor X", "Wolverine", "Phoenix", "Storm", "Shadowcat", "Magik"}
	n := []int{7, 4, 11, 8, 2, 19, 32, 3}

	// most straightforward
	sort.Ints(n)
	fmt.Println(n)

	// straight from the documentation; IntSlice -> Interface
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	// manual
	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	fmt.Println(sl)

	// other possibility
	// sort.StringSlice(sl).Sort()

	sort.Sort(sort.Reverse(sort.StringSlice(sl)))
	fmt.Println(sl)

	// could do it this way but why not take advantage of the type?
	// sort.Slice(studyGroup, func(i, j int) bool { return studyGroup[i] < studyGroup[j] })
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	// yay, can take advantage of the interface!
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}

// specific sort.Interface way
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
