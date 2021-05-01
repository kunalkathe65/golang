package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	d := newDeck()
	if(len(d) != 16){
		t.Errorf("Expected the deck size to be of 16 but got %v",len(d))
	}

	if(d[0] != "Ace Of Spades"){
		t.Errorf("Expected first card to be Ace Of Spades but got %v",d[0])
	}
	if(d[len(d) - 1] != "Four Of Clubs"){
		t.Errorf("Expected last card to be Four Of Clubs but got %v",d[len(d) - 1])
	}
}


func TestSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if(len(loadedDeck) != 16){
		t.Errorf("Expected the size of deck to be 16 but got %v",len(loadedDeck))
	}
	os.Remove("_decktesting")
}
