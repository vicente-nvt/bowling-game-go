package main

type Scorer struct {
	Throws []int
	frames []*Frame
}

func (s *Scorer) AddThrow(numberOfPinsDown int) {
	s.Throws = append(s.Throws, numberOfPinsDown)

	if len(s.frames) > 0 {
		currentFrame := s.frames[len(s.frames)-1]

		if !currentFrame.HasSecondThrow() {
			(*currentFrame).AddSecondThrow(numberOfPinsDown)
			return
		}

		scoreOfCurrentFrame, _ := currentFrame.GetFrameScore()
		s.startNewFrame(numberOfPinsDown, scoreOfCurrentFrame)
		(*currentFrame).AddTheFirstThrowOfNextFrame(numberOfPinsDown)

		return
	}

	s.startNewFrame(numberOfPinsDown, 0)
}

func (s *Scorer) startNewFrame(numberOfPinsDown int, scoreOfPriorFrame int) {
	frame := &Frame{}
	frame.AddFirstThrow(numberOfPinsDown)
	frame.AddPriorFrameScore(scoreOfPriorFrame)
	s.frames = append(s.frames, frame)
}

func (s *Scorer) GetFrameScore(index int) (int, error) {
	return s.frames[index].GetFrameScore()
}
