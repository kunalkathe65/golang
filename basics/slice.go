package main

import "fmt"

func main(){
	cards := []string{"Ace of Diamonds","Six Of Spade",newCard()}
	fmt.Println(cards)

	newCards := append(cards,"Eight of Club")
	fmt.Println(newCards)

	for i,card := range cards{
		fmt.Println(i,card)
	}
}

func newCard() string{
	return "Seven Of Club"
}
