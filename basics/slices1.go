package main

import (
  "fmt"
)

// show doubling of capacity as append() is used
func main() {
  // create a small slice and backing array
  x := make([]int, 1, 1)
  fmt.Printf("Length: %d, Capacity: %d\n", len(x), cap(x))
  for i := 1; i < 1000; i++ {
    x = append(x, i)
    fmt.Printf("Length: %d, Capacity: %d\n", len(x), cap(x))
  }
}
