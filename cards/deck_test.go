package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 20, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of clubs but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Spades" {
		t.Errorf("Expected Four of Spades but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewTestDeckFromFile(t *testing.T) {
	os.Remove("_deckTestFile")

	d := newDeck()
	d.saveToFile("_deckTestFile")

	loadFile := readFromFile("_deckTestFile")

	if len(loadFile) != len(d) {
		t.Errorf("File not matching!")
	}

	os.RemoveAll("_deckTestFile")
}
