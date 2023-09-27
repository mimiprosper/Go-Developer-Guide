package main

import "fmt"

var card string = "heart of gold" // variable declaration (global scope)

func main() {

	card := "ace of spade"	 // variable declaration (local scope) + initialization
	card = "queen of hearts"

	fmt.Println(card)
}