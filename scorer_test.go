package main

import "testing"

func TestShouldAddAThrowIntoScorer(t *testing.T) {
	numberOfPinsDown := 1

	scorer := &Scorer{}

	scorer.AddThrow(numberOfPinsDown)

	if scorer.Throws[0] != numberOfPinsDown {
		t.Errorf("The number of pins down for the first throw should be %v, but actually is %v",
			numberOfPinsDown, scorer.Throws[0])
	}
}

func TestShouldCalculateThePointsInAFrame(t *testing.T) {
	numberOfPinsDownInFirstThrow := 1
	numberOfPinsDownInSecondThrow := 4
	expectedScoreInFirstFrame := numberOfPinsDownInFirstThrow + numberOfPinsDownInSecondThrow
	scorer := &Scorer{}

	scorer.AddThrow(numberOfPinsDownInFirstThrow)
	scorer.AddThrow(numberOfPinsDownInSecondThrow)

	scoreOfFrame, _ := scorer.GetFrameScore(0)
	if scoreOfFrame != expectedScoreInFirstFrame {
		t.Errorf("The score of the first frame should be %v, but actually is %v",
			expectedScoreInFirstFrame, scoreOfFrame)
	}
}

func TestShouldCalculateThePointsInASecondFrameWithoutAMark(t *testing.T) {
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

	scoreOfFrame, _ := scorer.GetFrameScore(1)
	if scoreOfFrame != expectedScoreInSecondFrame {
		t.Errorf("The score of the second frame should be %v, but actually is %v",
			expectedScoreInSecondFrame, scoreOfFrame)
	}
}

func TestShouldNotCalculateThePointsOfASpareFrameUntilNextThrow(t *testing.T) {
	expectedMessage := "it isn't possible to calculate a spare frame without the next throw"
	scorer := &Scorer{}

	scorer.AddThrow(1)
	scorer.AddThrow(4)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(6)
	scorer.AddThrow(4)

	_, err := scorer.GetFrameScore(2)
	if err == nil || (err != nil && err.Error() != expectedMessage) {
		t.Errorf("There was expected an error with the following message:'%v', but it wasn't thrown", expectedMessage)
	}
}

func TestShouldCalculateThePointsOfASpareFrameAfterNextThrow(t *testing.T) {
	expectedScore := 29
	scorer := &Scorer{}

	scorer.AddThrow(1)
	scorer.AddThrow(4)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(6)
	scorer.AddThrow(4)
	scorer.AddThrow(5)

	obtainedScore, err := scorer.GetFrameScore(2)
	if err != nil {
		t.Errorf("There was an unexpected error: '%v'", err)
	}
	if expectedScore != obtainedScore {
		t.Errorf("The expected score was: '%v', but obtained was: '%v'", expectedScore, obtainedScore)
	}
}

func TestShouldCalculateThePointsOfASpareFrameAfterAStrike(t *testing.T) {
	expectedScore := 49
	scorer := &Scorer{}

	scorer.AddThrow(1)
	scorer.AddThrow(4)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(6)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(5)
	scorer.AddThrow(10)

	obtainedScore, err := scorer.GetFrameScore(3)
	if err != nil {
		t.Errorf("There was an unexpected error: '%v'", err)
	}
	if expectedScore != obtainedScore {
		t.Errorf("The expected score was: '%v', but obtained was: '%v'", expectedScore, obtainedScore)
	}
}

func TestShouldNotCalculateThePointsOfAStrikeWithoutNexFrameThrows(t *testing.T) {
	expectedMessage := "it isn't possible to calculate a strike frame without the throws of next frame"
	scorer := &Scorer{}

	scorer.AddThrow(1)
	scorer.AddThrow(4)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(6)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(5)
	scorer.AddThrow(10)
	scorer.AddThrow(0)

	_, err := scorer.GetFrameScore(4)
	if err == nil || (err != nil && err.Error() != expectedMessage) {
		t.Errorf("There was expected an error with the following message:'%v', but it wasn't thrown", expectedMessage)
	}
}

func TestShouldCalculateThePointsOfAStrikeFrameAfterNextThrow(t *testing.T) {
	expectedScore := 60
	scorer := &Scorer{}

	scorer.AddThrow(1)
	scorer.AddThrow(4)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(6)
	scorer.AddThrow(4)
	scorer.AddThrow(5)
	scorer.AddThrow(5)
	scorer.AddThrow(10)
	scorer.AddThrow(0)
	scorer.AddThrow(1)

	obtainedScore, err := scorer.GetFrameScore(4)
	if err != nil {
		t.Errorf("There was an unexpected error: '%v'", err)
	}
	if expectedScore != obtainedScore {
		t.Errorf("The expected score was: '%v', but obtained was: '%v'", expectedScore, obtainedScore)
	}
}
