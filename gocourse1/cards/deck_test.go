package main

import (
	"os"
	"testing"
)

// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, got %v", d[len(d)-1])
	}
}

func compareSlices(sla []string, slb []string) bool {
	if len(sla) != len(slb) {
		return false
	}

	for i := range sla {
		if sla[i] != slb[i] {
			return false
		}
	}

	return true
}

func TestShuffle(t *testing.T) {
	da := newDeck()
	da.Shuffle()
	db := newDeck()

	if compareSlices(da, db) {
		t.Errorf("Expected shuffle but shuffled deck is identical to new deck.")
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
