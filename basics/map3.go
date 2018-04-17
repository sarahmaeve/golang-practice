package main

import (
  "fmt"
)

// Create a map of student grades
// loop through and display all students with grades of B or better

func main() {
   grades := make(map[string]int)
   grades["Sarah"] = 95
   grades["Joe"] = 67
   grades["Teresa"] = 88
   grades["Adam"] = 96
   grades["Beth"] = 75

   for k, v :=  range(grades) {
     if v >= 80 {
       fmt.Println(k)
     }
   }
}
