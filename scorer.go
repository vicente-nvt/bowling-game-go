package main

type Scorer struct {
	Throws []int
	frames []*Frame
}

func (s *Scorer) AddThrow(numberOfPinsDown int) {
	s.Throws = append(s.Throws, numberOfPinsDown)

	if len(s.frames) > 0 {
		currentFrame := s.frames[len(s.frames)-1]
		shouldSetSecondThrow := !currentFrame.HasSecondThrow() && !currentFrame.IsStrike()

		if shouldSetSecondThrow {
			(*currentFrame).AddSecondThrow(numberOfPinsDown)

			hasPriorFrame := len(s.frames)-2 >= 0

			if hasPriorFrame {
				priorFrame := s.frames[len(s.frames)-2]
				priorFrame.AddTheSecondThrowOfNextFrame(numberOfPinsDown)
			}
			return
		}

		(*currentFrame).AddTheFirstThrowOfNextFrame(numberOfPinsDown)
		scoreOfCurrentFrame, _ := currentFrame.GetFrameScore()
		s.startNewFrame(numberOfPinsDown, scoreOfCurrentFrame)

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
