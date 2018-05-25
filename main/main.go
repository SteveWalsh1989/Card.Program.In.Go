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
 *
 * To Do:
 * 1. shuffle         : shuffle the deck
 * 2. saveToFile	  : save deck to file
 * 3. newDeckFromFile : load the deck from file
 *
 ******************************************/




func main() {

	cards := newDeck()  							 // create new deck called cards

	hand, remainingDeck := deal(cards, 5, ) // create a hand from deck


	hand.print()       								 // print all cards in the hand

	remainingDeck.print()						  	 // print all remaining cards


	}



