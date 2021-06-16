package main

type Frame struct {
	priorFrameScore int
	throws          []int
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

func (f *Frame) GetFrameScore() int {
	return f.priorFrameScore + f.throws[0] + f.throws[1]
}

func (f *Frame) AddPriorFrameScore(priorFrameScore int) {
	f.priorFrameScore = priorFrameScore
}

type Scorer struct {
	Throws []int
	frames []*Frame
}

func (s *Scorer) AddThrow(numberOfPinsDown int) {
	s.Throws = append(s.Throws, numberOfPinsDown)

	if len(s.frames) == 0 {
		frame := &Frame{}
		frame.AddFirstThrow(numberOfPinsDown)
		s.frames = append(s.frames, frame)
	} else {
		indexOfCurrentFrame := len(s.frames) - 1
		currentFrame := s.frames[indexOfCurrentFrame]
		if currentFrame.HasSecondThrow() {
			frame := &Frame{}
			frame.AddFirstThrow(numberOfPinsDown)
			frame.AddPriorFrameScore(currentFrame.GetFrameScore())
			s.frames = append(s.frames, frame)
		} else {
			(*currentFrame).AddSecondThrow(numberOfPinsDown)
		}
	}
}

func (s *Scorer) GetFrameScore(index int) int {
	return s.frames[index].GetFrameScore()
}
