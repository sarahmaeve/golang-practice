package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := newDeck()
	cards.Shuffle()
	cards.print()
}
