package main

import "fmt"

// code along with Udemy Advanced Go Programming

type Person struct {
	first, last string
}

func main() {
	// or use a bool instead if it has meaning
	people := make(map[Person]struct{})
	// struct{}{} : anonymous and empty
	people[Person{"Nina", "Simone"}] = struct{}{}

	_, ok := people[Person{"Nina", "Simone"}]
	fmt.Printf("%t\n", ok)
}
