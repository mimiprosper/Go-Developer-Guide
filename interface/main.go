package main

import "fmt"

// create interface type bot
type bot interface{
	getGreetings() string
}

// create type englishBot and spanishBot struct
type englishBot struct{}
type spanishBot struct{}

// bot function that accepts reciever ( ie b.printGreetings in OOP)
func printGreetings(b bot)  {
	fmt.Println(b.getGreetings())
}

// englishBot function that accepts reciever and returns a string
func (englishBot) getGreetings() string {
	return "hello there!!!"
}

// spanishBot function that accepts reciever and returns a string
func (spanishBot) getGreetings() string  {
	return "Hola espaniola"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}
