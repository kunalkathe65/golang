package main

import "fmt"

type deck []string

//Creates and returns a new Deck
func newDeck() deck{
	cards := deck{}
	cardSuites := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _,suit := range cardSuites{
		for _,value := range cardValues{
			cards = append(cards, value+" Of "+suit)
		}
	}
	return cards
}

//Creates and return d deal
func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

//Receiver functions
func  (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}
