package main


//-------------------
// Import Statements
//-------------------

import "testing"


//-------------------
//  Test Functions
//-------------------

func TestNewDeck(t *testing.T){

	d := newDeck()

	if len(d) != 52 { 										  // Scenario 1: Error, not correct deck size created

		t.Error("Expected deck lenght of: 52" +
			"		 \n Created length of: ", len(d)) 		  // print error
	}


	if d[0] != "Ace of Diamonds" {   						  // Scenario 2: Error, first card not correct


		t.Error("Expected first card: Ace of Diamonds" +
			"		 \n First card instead: ", d[0])          // print error

	}


	if d[len(d) -1 ] != "Queen of Clubs" { 					  // Scenario 3: Error, last card not correct


		t.Error("Expected first card: Queen of Clubs" +
			"		 \n First card instead: ", d[len(d) -1 ]) // print error

	}

}