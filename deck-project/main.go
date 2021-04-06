package main

func main(){
	cards := newDeck()
	inHand,remainingInDeck := deal(cards,5)
	inHand.print()
	remainingInDeck.print()
	cards.print()
}
