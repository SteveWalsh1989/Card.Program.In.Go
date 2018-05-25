package main


//-------------------
// Import Statements
//-------------------

import
(
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
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
 *  deal
 *
 *  deals out a hand of cards
 *
 * Input param  : d deck and handSize int
 *
 * Return param : first deck of size handSize, second deck is remaining cards
 */
func deal(d deck, handSize int) (deck, deck)  {

	return d[:handSize], d[handSize:]  // returns one deck from 0 - handSize and one with remaining elements


}

/*
 *  shuffle
 *
 *  receiver function for deck
 *
 *  shuffles deck by looping through slice
 *  generating random number between 0 and len(cards) - 1
 *  Swap the current card and the card at the cards[randomNumber]
 */
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) // seed for random number

	r := rand.New(source)  			      	        // generate new random number each time

	for i := range d {					      		// loop through deck

		newPos := r.Intn(len(d)-1)       			// generates random number between 0 and 1 less than length of slice

		d[i],d[newPos] = d[newPos], d[i] 			// swaps values of index i and new position newPos
	}

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
 *  newDeckFromFile
 *
 *  Locates file on disk and reads
 *
 *  returns deck
 */
func newDeckFromFile(fileName string) deck{

	bs,err :=  ioutil.ReadFile(fileName)     // bs: byteSlice, err: error object.

	if err != nil {  			             // Scenario 1: if error returned from ReadFile call

		fmt.Print("Error: ", err)	     //  log error

		os.Exit(1) 				     // quits program
	}
											 // Scenario 2: No error when reading file

	s := strings.Split(string(bs), ",") // convert bs to type string and use split func to return as slice of strings s

	return deck(s)  						 // return s after converted to type deck

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


