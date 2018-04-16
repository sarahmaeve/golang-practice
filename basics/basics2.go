package main
// simple hands on exercise #1
// from https://www.youtube.com/watch?v=fKJjQBlTHiA

import (
  "fmt"
)

func main() {
  x := 42
  y := "James Bond"
  z := true

  fmt.Println("x is", x)
  fmt.Println("y is", y)
  fmt.Println("z is", z)
  fmt.Printf("x is %d, y is %s, and z is %t\n", x, y, z )
}
