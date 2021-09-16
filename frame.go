package main

import "errors"

const numberOfPinsDownForASpare = 10
const numberOfPinsDownForAStrike = 10

type Frame struct {
	priorFrameScore          int
	throws                   []int
	firstThrowAfterSpare     int
	secondThrowAfterSpare    int
	hasFirstThrowAfterFrame  bool
	hasSecondThrowAfterFrame bool
}

func (f *Frame) addThrow(numberOfPinsDown int) {
	f.throws = append(f.throws, numberOfPinsDown)
}

func (f *Frame) AddFirstThrow(numberOfPinsDown int) {
	f.addThrow(numberOfPinsDown)
}

func (f *Frame) AddSecondThrow(numberOfPinsDown int) error {

	if !f.HasFirstThrow() {
		return errors.New("it isn't possible to add second throw before first throw")
	}

	f.addThrow(numberOfPinsDown)
	return nil
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

	return (f.throws[0] + f.throws[1]) == numberOfPinsDownForASpare
}

func (f *Frame) IsStrike() bool {

	if !f.HasFirstThrow() {
		return false
	}

	return f.throws[0] == numberOfPinsDownForAStrike
}

func (f *Frame) GetFrameScore() (int, error) {

	if f.IsSpare() && !f.hasFirstThrowAfterFrame {
		return 0, errors.New("it isn't possible to calculate a spare frame without the next throw")
	}

	if f.IsSpare() {
		return f.priorFrameScore + f.throws[0] + f.throws[1] + f.firstThrowAfterSpare, nil
	}

	if f.IsStrike() && (!f.hasFirstThrowAfterFrame || !f.hasSecondThrowAfterFrame) {
		return 0, errors.New("it isn't possible to calculate a strike frame without the throws of next frame")
	}

	if f.IsStrike() {
		return f.priorFrameScore + f.throws[0] + f.firstThrowAfterSpare + f.secondThrowAfterSpare, nil
	}

	if f.HasFirstThrow() && f.HasSecondThrow() {
		return f.priorFrameScore + f.throws[0] + f.throws[1], nil
	}

	return 0, nil
}

func (f *Frame) AddPriorFrameScore(priorFrameScore int) {
	f.priorFrameScore = priorFrameScore
}

func (f *Frame) AddTheFirstThrowOfNextFrame(firstThrowAfterFrame int) {
	f.firstThrowAfterSpare = firstThrowAfterFrame
	f.hasFirstThrowAfterFrame = true
}

func (f *Frame) AddTheSecondThrowOfNextFrame(secondThrowAfterFrame int) {
	f.secondThrowAfterSpare = secondThrowAfterFrame
	f.hasSecondThrowAfterFrame = true
}
