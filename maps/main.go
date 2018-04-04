package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
func main() {
	colors := map[string]string{
		"red":   "FF0000",
		"blue":  "00FF00",
		"green": "0000FF",
	}

	printMap(colors)
}
