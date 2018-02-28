package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	fmt.Println("Not Shuffled:")
	cards.print()
	cards.shuffle()
	fmt.Println("Shuffled:")
	cards.print()
}
