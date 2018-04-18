package main

import "fmt"

func rotateValuesLeft(valPtrs ...*int) {
	for x := 0; x < len(valPtrs)-1; x++ {
		// thank goodness for this feature which (I hope)
		// will avoid explicit buffering of a value
		*valPtrs[x], *valPtrs[x+1] = *valPtrs[x+1], *valPtrs[x]
	}
}

func rotateValuesRight(valPtrs ...*int) {
	for x := len(valPtrs) - 1; x > 0; x-- {
		*valPtrs[x], *valPtrs[x-1] = *valPtrs[x-1], *valPtrs[x]
	}
}

func main() {
	var a, b, c, d int = 1, 2, 3, 4
	rotateValuesLeft(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	rotateValuesRight(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	rotateValuesRight(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
}
