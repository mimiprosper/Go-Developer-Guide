package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a type of 'deck'
// which is a slice of strings
type deck []string

// reciever function ie acts like object of a class in OOP (d.print())
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// function that returns a deck of string
func newDeck() deck {

	cards := deck{}
	cardSuites := []string{"spades", "diamonds", "spades", "hearts", "kings"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// multiply return values and arguments
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to convert a deck to a slice of string
func (d deck) toString() string {
	// function to convert slice of string to string separated by commas
	return strings.Join([]string(d), ", ")
}

// save data to hard drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read file from hard drive + error handling. bs = byte of slice
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// print the error msg
		fmt.Println("Error:", err)
		// exit the program
		os.Exit(1)
	}
	// convert byte to slice of string. Note string(bs) = casting
	s := strings.Split(string(bs), ",")
	// convert slice of string to a deck
	return deck(s)
}

// shuffle a deck
func (d deck) shuffle()  {
	// generate random card sequence
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swapping position
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}

func main() {
	cards := newDeck()

	// save file to hard drive
	// cards.saveToFile("my_file")

	// open file of my hard drive. if not found the error is handled
	// cards := newDeckFromFile("my")
	cards.shuffle()
	cards.print()

}
