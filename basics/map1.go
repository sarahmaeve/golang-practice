package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"weight": 215,
		"height": 71, // trailing comma required for composite literals
	}

	fmt.Println(m["weight"])
	fmt.Println(m["width"]) // returns zero value due to missing key

	// comma, ok idiom
	if v, ok := m["width"]; ok {
		fmt.Println("This value is ok:", v)
	}
}
