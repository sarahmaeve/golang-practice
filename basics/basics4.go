package main
// https://www.youtube.com/watch?v=8TuXj6eiU2g
// using custom types
// https://www.youtube.com/watch?v=0c0cbqohzT0

import (
  "fmt"
)

type awesomenum int
var y int

func main() {
  var x awesomenum = 101
  // note : don't use new short assignment y := int(x) or y := x here.
  y = int(x)
  fmt.Printf("%T %v\n", y, y)
}
