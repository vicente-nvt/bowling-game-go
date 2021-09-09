package main

import (
	"testing"
)

func TestShouldAddFirstThrowToFrame(t *testing.T) {
	expectedLengthOfThrowsArray := 1
	numberOfPinsOfFirstThrow := 2
	frame := &Frame{}

	frame.AddFirstThrow(numberOfPinsOfFirstThrow)

	obtainedLengthOfThrowsArray := len(frame.throws)
	if expectedLengthOfThrowsArray != obtainedLengthOfThrowsArray {
		t.Errorf("The expected size of array should be '%v', but it actually is '%v'",
			expectedLengthOfThrowsArray, obtainedLengthOfThrowsArray)
	}

	obtainedValueOfFirstThrowInFrame := frame.throws[0]
	if numberOfPinsOfFirstThrow != obtainedValueOfFirstThrowInFrame {
		t.Errorf("The expected value of first throw should be '%v', but it actually is '%v'",
			numberOfPinsOfFirstThrow, obtainedValueOfFirstThrowInFrame)
	}
}

func TestShouldAddSecondThrowToFrame(t *testing.T) {
	expectedLengthOfThrowsArray := 2
	numberOfPinsOfFirstThrow := 3
	numberOfPinsOfSecondThrow := 5
	frame := &Frame{}
	frame.AddFirstThrow(numberOfPinsOfFirstThrow)

	frame.AddSecondThrow(numberOfPinsOfSecondThrow)

	obtainedLengthOfThrowsArray := len(frame.throws)
	if expectedLengthOfThrowsArray != obtainedLengthOfThrowsArray {
		t.Errorf("The expected size of array should be '%v', but it actually is '%v'",
			expectedLengthOfThrowsArray, obtainedLengthOfThrowsArray)
	}

	obtainedValueOfSecondThrowInFrame := frame.throws[1]
	if numberOfPinsOfSecondThrow != obtainedValueOfSecondThrowInFrame {
		t.Errorf("The expected value of second throw should be '%v', but it actually is '%v'",
			numberOfPinsOfSecondThrow, obtainedValueOfSecondThrowInFrame)
	}
}

func TestShouldIndicateThatFrameHasNoFirstThrowYet(t *testing.T) {

	frame := &Frame{}

	if frame.HasFirstThrow() {
		t.Errorf("It wasn't expected for the frame to have first throw yet")
	}
}

func TestShouldIndicateThatFrameHasFirstThrow(t *testing.T) {
	numberOfPinsDownInFirstThrow := 3
	frame := &Frame{}

	frame.AddFirstThrow(numberOfPinsDownInFirstThrow)

	if !frame.HasFirstThrow() {
		t.Errorf("It was expected for the frame to have first throw, but it hasn't")
	}
}

func TestShouldIndicateThatFrameHasNoSecondThrowYet(t *testing.T) {

	frame := &Frame{}

	if frame.HasFirstThrow() {
		t.Errorf("It wasn't expected for the frame to have second throw yet")
	}
}

func TestShouldIndicateThatFrameHasSecondThrow(t *testing.T) {
	numberOfPinsDownInFirstThrow := 3
	numberOfPinsDownInSecondThrow := 3
	frame := &Frame{}
	frame.AddFirstThrow(numberOfPinsDownInFirstThrow)

	frame.AddSecondThrow(numberOfPinsDownInSecondThrow)

	if !frame.HasFirstThrow() {
		t.Errorf("It was expected for the frame to have second throw, but it hasn't")
	}
}

func TestShouldNotAllowSecondThrowBeforeFirst(t *testing.T) {
	expectedErrorMessage := "it isn't possible to add second throw before first throw"
	numberOfPinsDownInSecondThrow := 3
	frame := &Frame{}

	err := frame.AddSecondThrow(numberOfPinsDownInSecondThrow)

	if err == nil {
		t.Errorf("It was expected an error with the following message: '%v'",
			expectedErrorMessage)
	}

	if err != nil && err.Error() != expectedErrorMessage {
		t.Errorf("It was expected an error with the following message: '%v', but the given message was: '%v'",
			expectedErrorMessage, err.Error())
	}
}

func TestShouldIndicateIfTheFrameIsNotSpareUntilTheSecondThrow(t *testing.T) {
	numberOfPinsDownInFirstThrow := 4
	frame := &Frame{}

	frame.AddFirstThrow(numberOfPinsDownInFirstThrow)

	if frame.IsSpare() {
		t.Errorf("It wasn't expected for the frame to be a Spare")
	}
}

func TestShouldIndicateThatTheFrameIsNotSpareWithout10PinsDown(t *testing.T) {
	numberOfPinsDownInFirstThrow := 4
	numberOfPinsDownInSecondThrow := 5
	frame := &Frame{}

	frame.AddFirstThrow(numberOfPinsDownInFirstThrow)
	frame.AddSecondThrow(numberOfPinsDownInSecondThrow)

	if frame.IsSpare() {
		t.Errorf("It wasn't expected for the frame to be a Spare")
	}
}

func TestShouldIndicateThatTheFrameIsASpare(t *testing.T) {
	numberOfPinsDownInFirstThrow := 3
	numberOfPinsDownInSecondThrow := 7
	frame := &Frame{}

	frame.AddFirstThrow(numberOfPinsDownInFirstThrow)
	frame.AddSecondThrow(numberOfPinsDownInSecondThrow)

	if !frame.IsSpare() {
		t.Errorf("It was expected for the frame to be a Spare")
	}
}
