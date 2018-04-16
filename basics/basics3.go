package main

import (
  "fmt"
)
// https://www.youtube.com/watch?v=PKbdA8LspD8
// exploring types
// hands on exercise 2 (modified)

var x int
var y string
var z bool

func main() {
  fmt.Printf("%T : %d, %T: '%s', %T: %t\n", x, x, y, y, z, z ) // show zero values for each type
}
