package main

import "fmt"
// from go tour; this seemed very go idiom

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
	  a, b = b, a + b  // similar to Python assignments for this problem
	  return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())  // read state of closure
	}
}
