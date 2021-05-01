package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"math/rand"
	"time"
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

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeckFromFile(filename string) deck{
	bs,err := ioutil.ReadFile(filename)
	if(err != nil){
		fmt.Println("Error occurred while reading",err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",")
	return deck(s)
}
