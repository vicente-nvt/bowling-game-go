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

	scoreOfFrame := scorer.GetFrameScore(0)
	if scoreOfFrame != expectedScoreInFirstFrame {
		t.Errorf("The score of the first frame should be %v, but actually is %v", expectedScoreInFirstFrame, scoreOfFrame)
	}
}

func TestShouldCalculateThePointsInASecondFrameWithourAMark(t *testing.T) {
	numberOfPinsDownInFirstThrow := 1
	numberOfPinsDownInSecondThrow := 4
	numberOfPinsDownInThirdThrow := 4
	numberOfPinsDownInFourthThrow := 5
	expectedScoreInSecondFrame := 14
	scorer := &Scorer{}

	scorer.AddThrow(numberOfPinsDownInFirstThrow)
	scorer.AddThrow(numberOfPinsDownInSecondThrow)
	scorer.AddThrow(numberOfPinsDownInThirdThrow)
	scorer.AddThrow(numberOfPinsDownInFourthThrow)

	scoreOfFrame := scorer.GetFrameScore(1)
	if scoreOfFrame != expectedScoreInSecondFrame {
		t.Errorf("The score of the second frame should be %v, but actually is %v", expectedScoreInSecondFrame, scoreOfFrame)
	}
}

func TestShouldValidateIfFramesExists(t *testing.T) {

}
