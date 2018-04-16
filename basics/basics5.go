package main
// simple 'bootcamp' program
// display all numbers from 1 - 100 divisible by 3

import (
  "fmt"
)

func main() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Println(i)
    }
  }
}
