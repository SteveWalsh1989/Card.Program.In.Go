package main

import "fmt"

/******************************************
 *
 * Cards Project
 *
 *
 *
 * Features:
 * 1. Create new deck of card
 * 2. Print all cards in a card to console
 * 3. deal			  : deal a hand of cards
 * 4. saveToFile	  : save deck to file
 * 5. newDeckFromFile : load the deck from file
 * 6. shuffle         : shuffle the deck
 *
 ******************************************/



func main() {

	cards := newDeck()									// create a new deck

	cards.shuffle() 									// shuffle the deck

	fmt.Println("\n//------------ \n Shuffled Deck " +
				   "\n//------------ \n")

	cards.print() 									 	// print the deck


	fmt.Println("\n//------------ \n Deck read from file " +
			       "\n//------------ \n")
	cards.saveToFile("newCardFile") 			// save deck to file


	cards2 := newDeckFromFile("newCardFile")   // load deck from file


	cards2.print() 										// print the deck


	fmt.Println("\n//------------ \n Hand dealt " +
				   "\n//------------ \n")

	h1,cards := deal(cards, 5)

	h1.print()


}



