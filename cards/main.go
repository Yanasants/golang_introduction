package main

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck() // the : is only used in the first declaration of the variable

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
