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
 * 3. deal			  : deal a hand of cards
 * 4. saveToFile	  : save deck to file
 * 5. newDeckFromFile : load the deck from file
 *
 * To Do:
 * 1. shuffle         : shuffle the deck
 *
 ******************************************/



func main() {


	cards := newDeck()

	cards.shuffle()

	cards.print()


	}



