// Straightforward interface assignments

package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{7.0}
	tri := triangle{height: 12.5, base: 5.75}

	printArea(sq)
	printArea(tri)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
