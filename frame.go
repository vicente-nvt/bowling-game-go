package main

import "errors"

type Frame struct {
	priorFrameScore         int
	throws                  []int
	firstThrowAfterSpare    int
	hasFirstThrowAfterFrame bool
}

func (f *Frame) addThrow(numberOfPinsDown int) {
	f.throws = append(f.throws, numberOfPinsDown)
}

func (f *Frame) AddFirstThrow(numberOfPinsDown int) {
	f.addThrow(numberOfPinsDown)
}

func (f *Frame) AddSecondThrow(numberOfPinsDown int) {
	f.addThrow(numberOfPinsDown)
}

func (f *Frame) HasFirstThrow() bool {
	return len(f.throws) > 0
}

func (f *Frame) HasSecondThrow() bool {
	return len(f.throws) == 2
}

func (f *Frame) IsSpare() bool {
	if !f.HasSecondThrow() {
		return false
	}

	return (f.throws[0] + f.throws[1]) == 10
}

func (f *Frame) GetFrameScore() (int, error) {

	if f.IsSpare() && !f.hasFirstThrowAfterFrame {
		return 0, errors.New("it isn't possible to calculate a spare frame without the next throw")
	}

	return f.priorFrameScore + f.throws[0] + f.throws[1] + f.firstThrowAfterSpare, nil
}

func (f *Frame) AddPriorFrameScore(priorFrameScore int) {
	f.priorFrameScore = priorFrameScore
}

func (f *Frame) AddFirstThrowAfterFrame(firstThrowAfterSpare int) {
	f.firstThrowAfterSpare = firstThrowAfterSpare
	f.hasFirstThrowAfterFrame = true
}
