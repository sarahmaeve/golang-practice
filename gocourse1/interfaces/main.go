package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	// mocking more complex logic
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

// interface types can't be receivers so no
// func (b bot) ...
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// unable to overload these functions in Go
// this is why we use interfaces instead.

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
//
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
