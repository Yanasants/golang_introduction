package main

func main() {
	//var card string = "Ace of Spades"
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
