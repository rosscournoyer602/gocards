package main

func main() {
	cards := newDeckFromFile("DeckDeck")
	cards.shuffle()
	cards.shuffle()
	cards.shuffle()
	cards.print()
}
