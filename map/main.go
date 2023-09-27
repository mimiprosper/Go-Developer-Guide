package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "ffffff"
	// colors["green"] = "aeff90"

	// // delete item in a map
	// delete(colors, "white")

	colors := map[string]string {
		"red" : "#ff0000",
		"green" : "#fr0988",
		"blue" : "#asff778",
		"white" : "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string)  {
	for color, hex := range c {
		fmt.Println(color, "has a hex:", hex)
	}
}