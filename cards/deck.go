package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), os.ModePerm)
}

func newDeckFromFile(filename string) deck {

	// get from file
	deckFromFile, err := ioutil.ReadFile(filename)
	// handle error
	if err != nil {

		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// split to slice
	cards := strings.Split(string(deckFromFile), ",")
	// convert to deck
	deckOfCards := deck(cards)

	return deckOfCards
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, card := range d {
		newPos := r.Intn(len(d))
		d[i] = d[newPos]
		d[newPos] = card
	}

	return d
}
