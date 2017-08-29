// testing if variable swapping works in Go

package main

import "fmt"
// Had to import math/big to match some of Python3 results for large targets
import "math/big"

func fib(n int) {
   a := big.NewInt(0)
   b := big.NewInt(1)
   i := 0
   for i < n + 1 {
     fmt.Println(a)
     // can't use + operator on pointers from NewInt
     a,b = b, a.Add(a,b)
     i += 1
   }
}

func main() {
  fib(1000)
}
