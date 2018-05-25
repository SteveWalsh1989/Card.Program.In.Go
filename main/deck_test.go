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

	if len(d) != 52 { 		// Scenario 1: Error, not correct deck size created

		t.Error("Expected deck lenght of: 52" +
			"		 \n Created length of: ", len(d)) // print error
	}

}