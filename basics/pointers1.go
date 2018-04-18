package main

import "fmt"

func square(p *int) {
	*p = *p * *p
}

func main() {
	x := 5
	fmt.Println(x, &x)
	square(&x)
	fmt.Println(x)
}
