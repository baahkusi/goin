package main

import "fmt"

func main() {
	cards := newDeck()

	print(cards.toString())

	filename := "deck.txt"

	shufs := cards.shuffle()

	shufs.saveDeckToFile(filename)

	loadedCards := newDeckFromFile(filename)

	print("\n\n")

	fmt.Println(loadedCards.toString())
}
