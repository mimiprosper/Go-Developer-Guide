package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

// function and return type - string
func newCard() string {
	return "diamond of fives"
}