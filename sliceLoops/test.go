// array is fixed list of things
// slice is an array that can grow or shrink

package main

import "fmt"

func main()  {
	// slice of string
	cards := [] string {"diamonds of five", newCard()}
	cards = append(cards, "heart of diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	
}

func newCard() string {
	return "heart of kings"
}