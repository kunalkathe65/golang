package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

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

//Creates and return the deal
func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

//Receiver functions
func  (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{
	data := []byte(d.toString())
	return ioutil.WriteFile(filename,data,0666)
}
