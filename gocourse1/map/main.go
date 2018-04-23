package main

import "fmt"

func printMap(m map[string]string) {
	for c, h := range m {
		fmt.Printf("Color: %v Hex Code: %v\n", c, h)
	}
}

func main() {
	// zero value of a map is an empty map -- no k,v pairs
	// var colors map[string]string

	// also no values using make
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	colors["white"] = "#ffffff"

	// how to delete a k,v pair
	// delete(colors, "white")
	printMap(colors)
}
