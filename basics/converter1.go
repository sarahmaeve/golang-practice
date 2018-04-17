// Go bootcamp exercise - Day 2
// Modify our miles to kilometers program to display in the following format:
// +-------------------------+
// | Miles: 50               |
// +-------------------------+
// | Kilometers: 80.47       |
// +-------------------------+
// Miles will be input by the user, and kilometers should be formatted to 2 decimal places.

package main

import (
  "fmt"
  "strings"
)

const MilestoKm float64 = 1.6039
var border string = "+" + strings.Repeat("-",30) + "+"

func main() {
  var miles float64
  fmt.Print("How many miles: ")
  fmt.Scanf("%f", &miles)
  fmt.Println(border)
  fmt.Printf("| Miles: %-21.2f |\n", miles)
  fmt.Println(border)
  fmt.Printf("| Kilometers: %-16.2f |\n", miles * MilestoKm)
  fmt.Println(border)
}
