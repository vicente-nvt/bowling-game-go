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

func TestShouldCalculateThePointsInAFrame(t *testing.T) {
	numberOfPinsDownInFirstThrow := 1
	numberOfPinsDownInSecondThrow := 4
	expectedScoreInFirstFrame := numberOfPinsDownInFirstThrow + numberOfPinsDownInSecondThrow
	scorer := &Scorer{}

	scorer.AddThrow(numberOfPinsDownInFirstThrow)
	scorer.AddThrow(numberOfPinsDownInSecondThrow)

	if scorer.Frames[0] != expectedScoreInFirstFrame {
		t.Errorf("The score of the first frame should be %v, but actually is %v", expectedScoreInFirstFrame, scorer.Frames[0])
	}
}

func TestShouldValidateIfFramesExists(t *testing.T) {

}
