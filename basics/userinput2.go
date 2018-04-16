package main
// miles to kilometers conversion with user input

import (
  "fmt"
)

func main() {
  var miles float64
  fmt.Print("How many miles: ")
  fmt.Scanf("%f", &miles) // pointer used as second argument
  fmt.Printf("%v miles is %v kilometers.\n", miles, miles * 1.6)
}
