package main

// test deck of slice
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// length of deck of slice
	if len(d) != 20 {
		t.Errorf("expected lenght of 20 but got %v", len(d))
		// t.Errorf("expected lenght of 20 but got", len(d))
	}

	// first item in the deck of slice
	if d[0] != "ace of spades" {
		t.Errorf("expected 1st item = ace of spaded but got", d[0])
	}

	// last item in the deck of slice
	if d[len(d)-1] != "four of lings" {
		t.Errorf("expected 1st item = ace of spaded but got", d[len(d)-1])
	}

}

// test file i/o
func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if loadedDeck != 16 {
		t.Errorf("expected 1st item = ace of spaded but got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
