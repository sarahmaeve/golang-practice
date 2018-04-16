package main

import (
  "fmt"
)
// https://www.youtube.com/watch?v=PKbdA8LspD8
// exploring types
// hands on exercise 2 (modified)
// follow up with https://www.youtube.com/watch?v=Q2EdLAL5vCA
// using Sprintf()

var x int = 42
var y string = "Natasha Romanoff"
var z bool = true

func main() {
  s := fmt.Sprintf("%T : %v, %T: '%v', %T: %v", x, x, y, y, z, z ) // show zero values for each type
  fmt.Println(s)
}
