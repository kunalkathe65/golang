package main

import "fmt"

type deck []string

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

//Receiver functions
func  (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}
