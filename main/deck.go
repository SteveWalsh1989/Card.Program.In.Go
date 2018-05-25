package main


//-------------------
// Import Statements
//-------------------

import
(
	"fmt"
	"strings"
	"io/ioutil"
)
//-------------------
//       Type
//-------------------

type deck []string 			               // create a new type: deck which is a slice of type: string



//-------------------
//     Functions
//-------------------

/*
 *  newDeck
 *
 *  creates new deck of cards
 *
 *  returns the deck of cards
 */
func newDeck() deck{

	cards := deck{}															// create new card of type deck

	cardSuits  := []string{"Diamonds", "Hearts", "Spades", "Clubs"}   		// new slice of strings for suits

	cardValues := []string{"Ace","Two","Three","Four","Five","Six","Seven", // new slice of strings for values
						   "Eight","Nine","Ten","Jack","King","Queen"}

	for _, suit := range  cardSuits { 										// loop through card suits

		for _, value := range cardValues{									// loop through card values

			cards = append(cards, value + " of " + suit)					// add the lists of card values to each suit
		}
	}
	return cards															// return deck of cards
}


/*
 *  print
 *
 *  receiver function for deck
 *
 *  prints all cards in the deck to console
 */
func (d deck) print() {

	for _, card := range d { 			// loop through cards in the deck

		fmt.Println(card) 				// print card
	}
}

/*
 *  toString
 *
 *  receiver function for deck
 *
 *  converts deck to string
 */
func (d deck) toString() string {

	return strings.Join([]string(d), ",") // use string package join func to link elements of []string converted deck by comma
}



/*
 *  saveToFile
 *
 *  receiver function for deck
 *
 *  Saves string to file
 *
 *  Input param: fileName
 *
 *  Returns : can return error
 */
func (d deck) saveToFile(fileName string) error{

	return ioutil.WriteFile(fileName,[]byte(d.toString()), 0666) // coverts deck to string then byte slice and saves to file
																// 0666: Anyone can read/write to the file

}


/*
 *  deal
 *
 *  deals out a hand of cards
 *
 * Input param  : d deck and handSize int
 *
 * Return param : two decks
 */
func deal(d deck, handSize int) (deck, deck)  {

	return d[:handSize], d[handSize:]  // returns one deck from 0 - handSize and one with remaining elements


}


