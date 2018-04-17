package main
// Errors exercise from https://tour.golang.org/methods/20

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
   e := ErrNegativeSqrt(x)
   return x, e
  }

  z := x / 2
  for i := 1; i <= 10; i++ {
	z -= (z*z - x) / (2*x)
  }
  return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
