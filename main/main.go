package main


/******************************************
 *
 * Cards Project
 *
 * --- Runs from terminal but not in GoLand ---
 *
 *
 * Features:
 * 1. Create new deck of card
 * 2. Print all cards in a card to console
 *
 *
 * To Do:
 * 1. shuffle         : shuffle the deck
 * 2. deal			  : deal a hand of cards
 * 3. saveToFile	  : save deck to file
 * 4. newDeckFromFile : load the deck from file
 *
 ******************************************/




func main() {

	cards := newDeck()  // create new deck called cards

	cards.print()       // print all cards

	}



