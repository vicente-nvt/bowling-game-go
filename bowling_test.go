package main

import "testing"

func TestShouldAddAThrowIntoScorer(t *testing.T) {
	numberOfPinsDown := 1
	scorer := &Scorer{}

	scorer.AddThrow(numberOfPinsDown)

	if scorer.Throws[0] != numberOfPinsDown {
		t.Errorf("The number of pins down for the first throw should be %v, but actually is %v", numberOfPinsDown, scorer.Throws[0])
	}
}
